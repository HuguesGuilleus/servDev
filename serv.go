package main

import (
	"errors"
	"flag"
	"fmt"
	"html"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
)

var (
	noMinify     = false
	noLivereload = false
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of serv: [OPTIONS] rootPath\n")
		flag.PrintDefaults()
	}
	address := flag.String("a", ":8000", "Listen address")
	flag.BoolVar(&noMinify, "m", false, "No minify served file (css and js)")
	flag.BoolVar(&noLivereload, "l", false, "No inject hot reload code in HTML")
	flag.Parse()

	go watching()
	os.Stdout.WriteString("\033[H\033[2J")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r, http.Dir(flag.Arg(0)))
	})
	http.HandleFunc("/$change", handlerChange)

	log.SetFlags(log.Ltime)
	log.Println("[listen]", *address, flag.Arg(0))
	log.Fatal("[listen-error]", http.ListenAndServe(*address, nil))
}

func handlerChange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "text/event-stream")

	flusher, ok := w.(http.Flusher)
	if !ok {
		return
	}

	w.Write([]byte("retry: 50\n\n"))
	flusher.Flush()

	lastEventCond.L.Lock()
	defer lastEventCond.L.Unlock()
	lastEventCond.Wait()

	w.Write([]byte("data: change\n\n"))
}

var lastEventCond = sync.NewCond(&sync.Mutex{})

func watching() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	var indexer func(string)
	indexer = func(root string) {
		watcher.Add(root)
		entrys, _ := os.ReadDir(root)
		for _, entry := range entrys {
			if entry.IsDir() {
				indexer(filepath.Join(root, entry.Name()))
			}
		}
	}
	root := filepath.Join(flag.Arg(0))
	go indexer(root)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&(fsnotify.Create|fsnotify.Rename) != 0 {
				go indexer(event.Name)
			}
			lastEventCond.Broadcast()
			logRequest(nil, "fsnotif", event.Op.String()+" "+event.Name)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			logRequest(nil, "error.fsnotif", err.Error())
		}
	}
}

func handle(w http.ResponseWriter, r *http.Request, fsys http.FileSystem) {
	w.Header().Add("Server", "servHttp/2")
	p := path.Clean(r.URL.Path)

	if strings.HasSuffix(p, "/index.html") {
		logRequest(r, "redirect.noindex", "./")
		http.Redirect(w, r, "./", http.StatusTemporaryRedirect)
		return
	}

	file, info, err := open(fsys, p)
	if err != nil {
		if p == "/favicon.ico" {
			logRequest(r, "favicon", p)
			w.Header().Set("Content-type", "image/webp")
			w.Write([]byte(favicon))
		} else {
			responseError(w, r, err)
		}
		return
	}
	defer file.Close()

	if isDir, slashSuffix := info.IsDir(), strings.HasSuffix(r.URL.Path, "/"); isDir && !slashSuffix {
		p += "/"
		logRequest(r, "redirect.toslash", p)
		http.Redirect(w, r, p, http.StatusTemporaryRedirect)
	} else if !isDir && slashSuffix {
		logRequest(r, "redirect.noslash", p)
		http.Redirect(w, r, p, http.StatusTemporaryRedirect)
	} else if isDir {
		if indexFile, indexInfo, err := open(fsys, path.Join(p, "index.html")); err == nil {
			defer indexFile.Close()
			logRequest(r, "index.file", "")
			serveContent(w, r, indexInfo.Name(), indexInfo.ModTime(), indexFile)
		} else {
			childs, err := file.Readdir(0)
			if err != nil {
				responseError(w, r, err)
				return
			}
			responseIndex(w, r, childs)
		}
	} else {
		logRequest(r, "file", "")
		serveContent(w, r, info.Name(), info.ModTime(), file)
	}
}

func open(fsys http.FileSystem, path string) (file http.File, info fs.FileInfo, err error) {
	file, err = fsys.Open(path)
	if err != nil {
		return
	}

	info, err = file.Stat()
	if err != nil {
		file.Close()
	}

	return
}

func serveContent(w http.ResponseWriter, r *http.Request, name string, modtime time.Time, content io.ReadSeeker) {
	switch {
	case strings.HasSuffix(name, ".html"):
		if noLivereload {
			http.ServeContent(w, r, name, modtime, content)
		} else {
			data, err := io.ReadAll(content)
			if err != nil {
				responseError(w, r, err)
				return
			}
			w.Write(htmlHeadRegexp.ReplaceAll(data, []byte(`<head$1>`+eventSourceScript)))
		}

	case strings.HasSuffix(name, ".js") || name == "js":
		if noMinify {
			http.ServeContent(w, r, name, modtime, content)
		} else {
			w.Header().Set("Content-Type", "application/javascript")
			if err := js.Minify(nil, w, content, nil); err != nil {
				log.Printf("[error.minify] %q: %v", name, err)
			}
		}

	case strings.HasSuffix(name, ".css") || name == "css":
		if noMinify {
			http.ServeContent(w, r, name, modtime, content)
		} else {
			w.Header().Set("Content-Type", "text/css")
			if err := css.Minify(nil, w, content, nil); err != nil {
				log.Printf("[error.minify] %q: %v", name, err)
			}
		}

	default:
		http.ServeContent(w, r, name, modtime, content)
	}
}

func responseError(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, fs.ErrNotExist) {
		logRequest(r, "noFound", err.Error())
		w.WriteHeader(http.StatusNotFound)
		beginTemplate(w, r, "Error: File not found")
	} else {
		logRequest(r, "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		beginTemplate(w, r, "Error: internal error")
	}

	w.Write([]byte(`<pre id=err>`))
	w.Write([]byte(html.EscapeString(err.Error())))
}

func responseIndex(w http.ResponseWriter, r *http.Request, childs []fs.FileInfo) {
	logRequest(r, "index.auto", "")
	beginTemplate(w, r, "Index")

	sort.Slice(childs, func(i, j int) bool {
		if childs[i].IsDir() == childs[j].IsDir() {
			return childs[i].Name() < childs[j].Name()
		} else {
			return childs[i].IsDir()
		}
	})
	readme := ""
	w.Write([]byte(`<div id=files>`))
	for _, file := range childs {
		name := html.EscapeString(file.Name())
		if file.IsDir() {
			w.Write([]byte(`<span></span><span>-----</span>`))
			w.Write([]byte(`<a href="`))
			w.Write([]byte(name))
			w.Write([]byte(`/">`))
			w.Write([]byte(name))
			w.Write([]byte(`/</a>`))
		} else {
			w.Write([]byte(`<time datetime="`))
			w.Write([]byte(file.ModTime().Format(time.RFC3339)))
			w.Write([]byte(`">`))
			w.Write([]byte(file.ModTime().String()))
			w.Write([]byte(`</time>`))

			w.Write([]byte(`<span class=size>`))
			w.Write([]byte(int2string(int(file.Size()))))
			w.Write([]byte(` o</span>`))

			w.Write([]byte(`<a href="`))
			w.Write([]byte(name))
			w.Write([]byte(`">`))
			w.Write([]byte(name))
			w.Write([]byte(`</a>`))

			switch strings.ToLower(name) {
			case "readme", "readme.markdown", "readme.md", "readme.txt":
				readme = file.Name()
			}
		}
	}
	w.Write([]byte(`</div>`))

	// Print readme + Search + time use user local format
	w.Write([]byte(`<script>((d,H,P,T,`))
	defer w.Write([]byte(`})(document,"hidden","previousElementSibling","innerText")</script>`))
	w.Write([]byte(`c=(t,e=d.createElement(t),_=d.body.append(e))=>e,`))
	w.Write([]byte(`q=(s,f)=>d.querySelectorAll(s).forEach(f),`))
	w.Write([]byte(`l=t=>t.toLowerCase()`))
	w.Write([]byte(`)=>{`))
	if readme != "" {
		w.Write([]byte(`fetch("`))
		w.Write([]byte(readme))
		w.Write([]byte(`").then(r=>r.text()).then(t=>{c("hr");c("pre")[T]=t});`))
	}
	w.Write([]byte(`q("time",t=>t[T]=new Date(t.dateTime).toLocaleString());`))
	w.Write([]byte(`q("input",i=>{i[H]=0;i.focus();i.oninput=_=>q("#files>a",a=>a[P][P][H]=a[P][H]=a[H]=!l(a[T]).includes(l(i.value)))})`))
}

func int2string(i int) string {
	if i < 1000 {
		return strconv.Itoa(i % 1000)
	}
	return int2string(i/1000) + fmt.Sprintf(" %03d", i%1000)
}

func logRequest(r *http.Request, operation string, arg string) {
	path := ""
	if r != nil {
		path = r.URL.Path
	}
	if arg == "" {
		log.Printf("[%10s] %s", operation, path)
	} else {
		log.Printf("[%10s] %s :: %s", operation, path, arg)
	}
}

func beginTemplate(w http.ResponseWriter, r *http.Request, title string) {
	title = html.EscapeString(title)

	w.Write([]byte(`<!DOCTYPE html><html lang=en><head><meta charset=utf-8>`))
	w.Write([]byte(`<meta name=viewport content="width=device-width,initial-scale=1">`))
	w.Write([]byte(eventSourceScript))
	w.Write([]byte(`<style>`))
	w.Write([]byte(`body{margin:1em;font-family:monospace;font-size:xx-large;background:#eae5dc;color:black}`))
	w.Write([]byte(`pre{white-space:pre-wrap;color:#0009}`))
	w.Write([]byte(`input{all:unset;border:unset;font:unset;background:none;color:#0065c8}`))
	w.Write([]byte(`h1,.p,input{display:table;margin:0 0 0.5ex;padding:0.2em 0.5em;background:white}`))
	w.Write([]byte(`nav,#files{display:grid;grid-template-columns:auto 1fr;grid-gap:0 1ex}`))
	w.Write([]byte(`a{color:#0065c8;text-decoration:none}`))
	w.Write([]byte(`a:hover{color:#00008b;text-decoration:underline}`))
	w.Write([]byte(`#files{grid-template-columns:auto auto 1fr}`))
	w.Write([]byte(`time,span{color:#0009;text-align:right}`))
	w.Write([]byte(`time::after{content:"]"}time::before{content:"["}`))
	w.Write([]byte(`hr{border:0.17ex solid}`))
	w.Write([]byte(`*[hidden]{display:none}`))
	w.Write([]byte(`</style><title>`))
	w.Write([]byte(title))
	w.Write([]byte(`</title></head><body><h1>`))
	w.Write([]byte(title))
	w.Write([]byte(`</h1>`))

	// Path
	w.Write([]byte(`<nav>`))
	w.Write([]byte(`<div class=p>`))
	splitedPath := strings.Split(r.URL.Path, "/")
	pathLastElement := html.EscapeString(splitedPath[len(splitedPath)-1])
	splitedPath = splitedPath[:len(splitedPath)-1]
	all := strings.Builder{}
	for _, p := range splitedPath {
		p = html.EscapeString(p)
		all.WriteString(p)
		all.WriteByte('/')
		fmt.Fprintf(w, `<a href="%s">%s/</a>`, all.String(), p)
	}
	if pathLastElement != "" {
		w.Write([]byte(`<a>`))
		w.Write([]byte(pathLastElement))
		w.Write([]byte(`</a>`))
	}
	w.Write([]byte(`</div>`))

	w.Write([]byte(`<input type=search placeholder=Search hidden></nav>`))
}

var htmlHeadRegexp = regexp.MustCompile(`<head([^>]*)>`)

const eventSourceScript = `<script>new(EventSource||Array)("/$change").onmessage=_=>location.reload()</script>`

// Image from GitHub server icon.
const favicon = "RIFF\b\x04\x00\x00WEBPVP8X\n\x00\x00\x00\x10\x00\x00\x00\xfb\x00\x00\xfb\x00\x00ALPHA\x03\x00\x00\x01\x90l۶i;\xebٶm;\xb6\x93\xaam۶\xf1\x01\xb6\x93\x9am۶m]\xee\xe2n\xc9Zm\x9e\xe2\x89\bGn\x1b9\x12\xb59\x1fgf\xbb>@\xa6\xbbydV4\x02WE\xa6\xbbP\x993\xceZ\x15\xc0\xacg\xa6gʓ\xb4֩p\xb6#C\xd8g/\xb6*\xa8Y\x17y\b\x12zH\xc1\xedx\x84\x18\xd1\x0f\x14\xe0\xeeG\t\xe1}ZA\ued17\f\xab\x15薋P[\xc1\xae\x96\x04\xa7p;\xeb¯\xa5\x02^3~ۑ\xdb\xc2\xce\xfd;r\xdfݹ\x95+\xe8\x95r\xeb\x84]\an\xa3\xb0\x1b\xc1m\x8a\xd6\xc6y\xba\xb4\x9e\xcd\xd3u\x92\xfdg\xff\xb3y\xba6jM\xe16M\xab\x1e\xe9\xd2:\"\xeeg\xff3\xd2UOk\x9a9\xaa\x03\xba.\x1a\x8d\xff\x9dY\xa1\t\nz\xe3\xb9\rĮ?\xb7\xe6\xd85\xe1\x16\xe3D\xceɿ\x86\xfb&r\u05c8\xddd\xe8\xaf\x140\xb0\xf0\x13\xb7\x1f\x12\xc6\x16\xa6\xe36\x85\x04\bx\x83\xda+\u007f\x92PͿ\x98Y뒌\xbaa֛\xa44؆\x97m\x10ɩ\xf1g\xb4>5\"I\xc5-\xb7#e_\x16K\xc2\xca[\xfb\x13\xe6\xe6\xc4\xda<\x12\x98o\x93Ikw\x1f\x00\xd7\ued53\x1a\xfb\x90yq\x01u{\x8e\x85WϺ\xfe\"y\xf7>\x06\xf2H\xa9\xedh/oi\\[?V@{\xd1\xc7M\x94\xf8\v\nl\xe7\xe2\x05\xa9\xf6F\xc1\xedu\x95\x9c\xf3o\x05\xb8ߕBľT\x90{\x13/\x82\xdb\x05\x05\xba\xb3\xae\x12tW\xb0\xeb,a{\xe2\x19n/|\xf8\xf5V\xc0\xeb\xc9\xef8r\x87\xf9\xf7\x1d\xac\xc8Y\xfc\xb9\xd5UЫͭ\x17v=\xb8\x8d\xc5n4\xb7\xa9Z#\x1b\xe9Һ\xdaH\xd7Z\xf6\x9f\xcdm\xaa\xb1\xcfk\x98\xfe3\x19a\xec\xff}\x1b\xfb\x9a\ac_\xebb\xeck\x9c\f}m\x9b\x1fqw\f\xb9C\x86\xbe\x96\xb5\x87\x91\xafa~\xe2E\xfcuŭ\xa3\x88Ѭs\xa8\x9dv%\tż\xc4쵔\x11\xad*\xcc3J\x15$\xa5\x92\xe7x\xbd\xaa$9ŝ\x85\xfb\x8eX\x92\x94K\xebGH=\xef\xe3&n\x06\xb9\xc7a+\xc8\r\xe8C=\xbcHb~\xb5\xbb\x8d\x9d5\x0fZ\xb3\xc6v\xab\xedG\xe6\xc5\xf97\x9f\xban\xcf\x01p\xedY7\xa5\x99\x9fD\x05\x1b~+\x90\xfdZ_ M\xc2j\x87\x02\x9ace\xbc(M\xbe(\xb0}o%\xc8P\xbb\x82\x9bm\xb0\x18\xdd\x15\xe4\xfa\bQ˂\x99\xb5\x9e\b\x81o\x15\xe8^\xfbK0S\xc1n\x9a\x84\u007f\xbe\u007f\xe2\xf6#\x8a\xdf\x14\x05\xbc\t\xfcn!w\x9d\xff\xbeB\xce\x19ŭ\xb9\x82^\x13n\x03\xb1\xeb\xcfm\x02v\xe3\x8d}^\xc3\xf4\x9fi\x9c)ZK\xc6\xea\xd2z4V\xd7\x01\xf6\x9f\xcdm\n\xb7Q\nz#\xb8u®\x03\xb7r\xecJ\xb8\xb9\u007fC\xee\xab;q\xb7\r\xb9\xcdĮ\x05r͈\xbf\x93\xb8\x9dq\x11\xa0\x9a\x13\xb6\x9a$\xa1\x95\xa8-%\x11<\x8favʋd\x14u\x1f\xb1\xfbQ$\xa5Ѓx\x1d\x8b 9\xb9/\xb0`e\x99\xe7N\xa2J\\\x02t\x18չ5\x8dĕ>\xe5\x14ė[NNN#\x99\xb9\xa7\x977\x02Wy\x9a;\x99\xedF\x00VP8 \xa0\x00\x00\x00\xf0\x10\x00\x9d\x01*\xfc\x00\xfc\x00>\x91H\xa1M%\xa4#\" (\x00\xb0\x12\tin\xe1v\xb1\x1b@\t\xec\x03\xdfl\x9c\x87\xbe\xd99\x0f}\xb2r\x1e\xfbd\xe4=\xf6\xc9\xc8{퓐\xf7\xdb'!\xef\xb6NC\xdfl\x9c\x87\xbe\xd99\x0f}\xb2r\x1e\xfbd\xe4=\xf6\xc9\xc8{퓐\xf7\xdb'!\xef\xb6NC\xdfl\x9c\x87\xbe\xd99\x0f}\xb2r\x1e\xfbd\xe4=\xf6\xc9\xc8{퓐\xf7\xdb'!\xef\xb6NC\xdfl\x9c\x87\xbe\xd99\x0f}\xb2r\x1e\xfbd\xe1@\x00\xfe\xff\xbc\x11|@\x00\x00\x00\x00\x00\x00\x00\x00\x00"

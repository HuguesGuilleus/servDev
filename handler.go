// HTTP Server for developping project
// 2019, 2020 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"./templ"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
)

// The server contain the information.
type server struct {
	Dir http.Dir
}

func (s *server) Open(p string) (http.File, os.FileInfo, error) {
	f, err := s.Dir.Open(p)
	if err != nil {
		return nil, nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, nil, err
	}
	return f, stat, nil
}

func serverSignature(w http.ResponseWriter) {
	w.Header().Add("Server", "servHttp/1")
}

// Serve the 'favicon.ico' file.
func (s *server) favicon(w http.ResponseWriter, r *http.Request) {
	serverSignature(w)
	icon, err := s.Dir.Open(path.Join(string(s.Dir), "favicon.ico"))
	if err != nil {
		handleFavicon(w, r)
		return
	}
	defer icon.Close()
	w.Header().Add("Content-Type", "image/x-icon")
	io.Copy(w, icon)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serverSignature(w)

	p := r.URL.Path
	file, stat, err := s.Open(p)
	if err != nil {
		handleNotFound(w, r)
		return
	}
	defer file.Close()

	if stat.IsDir() {
		if !strings.HasSuffix(p, "/") {
			http.Redirect(w, r, p+"/", http.StatusMovedPermanently)
			return
		}
		if index, stat, err := s.Open(p + "index.html"); err == nil {
			defer index.Close()
			handleFile(w, r, index, stat)
			return
		}
		handleIndex(w, file, p)
	} else {
		handleFile(w, r, file, stat)
	}
}

// handleFile write the file into the http.ResponseWriter.
func handleFile(w http.ResponseWriter, r *http.Request, file http.File, stat os.FileInfo) {
	log.Println("[FILE]", r.URL.Path)
	http.ServeContent(w, r, stat.Name(), stat.ModTime(), file)
}

// handleIndex generate the index of the directory into the http.ResponseWriter.
func handleIndex(w http.ResponseWriter, dir http.File, p string) {
	files, err := dir.Readdir(-1)
	if err != nil {
		http.Error(w, fmt.Sprintf("Read dir %q fail: %v\r\n", p, err), http.StatusInternalServerError)
		return
	}
	sort.Slice(files, func(i, j int) bool {
		if files[i].IsDir() == files[j].IsDir() {
			return files[i].Name() < files[j].Name()
		} else {
			return files[i].IsDir()
		}
	})
	log.Println("[INDEX]", p)
	w.Header().Add("Content-type", "text/html; charset=utf-8")
	templ.Dir.Execute(w, files)
}

// ResponseWriter response with status not found.
func handleNotFound(w http.ResponseWriter, r *http.Request) {
	log.Println("[NOT FOUND]", r.URL.Path)
	w.Header().Add("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	templ.NotFound.Execute(w, strings.Split(r.URL.Path, "/"))
}

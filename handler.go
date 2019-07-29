// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main
import (
	"io"
	"log"
	"os"
	"net/http"
	"./templ"
	"strings"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	// Get information
	url := r.URL.Path
	file,err1 := os.Open("."+url)
	if err1 != nil {
		log.Print("[Error] ",url)
		handleNotFound(w,r)
		return
	}
	defer file.Close()
	info,err2 := file.Stat()
	if err2 != nil {
		log.Print("[Error] ",url)
		handleNotFound(w,r)
		return
	}
	// Response
	if info.IsDir() {
		if url[len(url)-1] == '/' {
			if handleIndex(w,url) {
				log.Print("[Index] ",url)
			} else {
				log.Print("[Directory] ",url)
				handleDir(w,file)
			}
		} else {
			log.Print("[Redirect directory] ",url)
			handleRedirect(w,url+"/")
		}
	} else {
		log.Print("[File] ",url)
		handleFile(w,file)
	}
}

func handleFile(w http.ResponseWriter, file *os.File) {
	w.Header().Add("Content-type",mime(file.Name()))
	io.Copy(w,file)
}

func handleIndex(w http.ResponseWriter, path string) (ok bool) {
	index,err := os.Open("."+path+"index.html")
	if err != nil {
		return false
	}
	defer index.Close()
	io.Copy(w,index)
	return true
}

func handleDir(w http.ResponseWriter, dir *os.File) {
	w.Header().Add("Content-type","text/html; charset=utf-8")
	files,_ := dir.Readdir(0)
	templ.Dir.Execute(w,files)
}

func handleRedirect(w http.ResponseWriter, url string) {
	w.Header().Add("Content-type","text/html; charset=utf-8")
	w.Header().Add("Location",url)
	w.WriteHeader(301)
	templ.Redirect.Execute(w,url)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type","text/html; charset=utf-8")
	w.WriteHeader(404)
	templ.NotFound.Execute(w,strings.Split(r.URL.Path,"/"))
}

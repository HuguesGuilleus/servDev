// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"./templ"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// func handleFavicon(w http.ResponseWriter, _ *http.Request) {
// 	w.Header().Add("Content-type", "image/svg+xml")
// 	log.Println("[Favicon]")
// 	w.Write([]byte(`<?xml version="1.0"?><svg xmlns="http://www.w3.org/2000/svg" width="256" height="256" class="octicon octicon-device-desktop" viewBox="0 0 16 16" version="1.1" aria-hidden="true"><path fill-rule="evenodd" d="M15 2H1c-.55 0-1 .45-1 1v9c0 .55.45 1 1 1h5.34c-.25.61-.86 1.39-2.34 2h8c-1.48-.61-2.09-1.39-2.34-2H15c.55 0 1-.45 1-1V3c0-.55-.45-1-1-1zm0 9H1V3h14v8z"></path></svg>`))
// }

func handleMain(w http.ResponseWriter, r *http.Request) {
	// Get information
	url := r.URL.Path
	file, err1 := os.Open("." + url)
	if err1 != nil {
		log.Print("[Error] ", url)
		handleNotFound(w, r)
		return
	}
	defer file.Close()
	info, err2 := file.Stat()
	if err2 != nil {
		log.Print("[Error] ", url)
		handleNotFound(w, r)
		return
	}
	// Response
	if info.IsDir() {
		if url[len(url)-1] == '/' {
			if handleIndex(w, url) {
				log.Print("[Index] ", url)
			} else {
				log.Print("[Directory] ", url)
				handleDir(w, file)
			}
		} else {
			log.Print("[Redirect directory] ", url)
			handleRedirect(w, url+"/")
		}
	} else {
		log.Print("[File] ", url)
		handleFile(w, file)
	}
}

func handleFile(w http.ResponseWriter, file *os.File) {
	w.Header().Add("Content-type", mime(file.Name()))
	io.Copy(w, file)
}

func handleIndex(w http.ResponseWriter, path string) (ok bool) {
	index, err := os.Open("." + path + "index.html")
	if err != nil {
		return false
	}
	defer index.Close()
	io.Copy(w, index)
	return true
}

func handleDir(w http.ResponseWriter, dir *os.File) {
	w.Header().Add("Content-type", "text/html; charset=utf-8")
	files, _ := dir.Readdir(0)
	templ.Dir.Execute(w, files)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(404)
	templ.NotFound.Execute(w, strings.Split(r.URL.Path, "/"))
}

func handleRedirect(w http.ResponseWriter, url string) {
	w.Header().Add("Content-type", "text/html; charset=utf-8")
	w.Header().Add("Location", url)
	w.WriteHeader(301)
	templ.Redirect.Execute(w, url)
}

// REDIRECTOIN SERVER

// redirect contain a string of the protocole and the server and port
type redirect string

func (dest redirect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u := string(dest) + r.URL.Path
	log.Println("[REDIRECT SERV]", u)
	handleRedirect(w, u)
}

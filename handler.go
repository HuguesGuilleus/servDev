// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"./templ"
	"io"
	"log"
	"net/http"
	"path"
	"strings"
)

// The server contain the information.
type server struct {
	Dir http.Dir
}

// Serve the 'favicon.ico' file.
func (s *server) favicon(w http.ResponseWriter, r *http.Request) {
	icon, err := s.Dir.Open(path.Join(string(s.Dir), "favicon.ico"))
	if err != nil {
		handleFavicon(w, r)
		return
	}
	w.Header().Add("Content-type", "image/x-icon")
	io.Copy(w, icon)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	defer func() {
		if err := recover(); err != nil {
			log.Println("[ERROR]", p, "::", err)
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Internal error")
		}
	}()

	file, err := s.Dir.Open(p)
	if err != nil {
		handleNotFound(w, r)
		return
	}

	if isDir(file) {
		if !strings.HasSuffix(p, "/") {
			http.Redirect(w, r, p+"/", http.StatusMovedPermanently)
			return
		}
		// if index, err := s.Dir.Open(p + "index.html"); err == nil {
		if index, err := s.Dir.Open(path.Join(p, "index.html")); err == nil {
			handleFile(w, index, p, "index.html")
			return
		}
		handleIndex(w, file, p)
	} else {
		handleFile(w, file, p, p)
	}
}

// Return true if the file is a directory. In error case, the fonction panic.
func isDir(file http.File) bool {
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	return info.IsDir()
}

// handleFile write the file into the http.ResponseWriter.
func handleFile(w http.ResponseWriter, file http.File, p, name string) {
	log.Println("[FILE]", p)
	w.Header().Add("Content-type", mime(name))
	io.Copy(w, file)
}

// handleIndex generate the index of the directory into the http.ResponseWriter.
func handleIndex(w http.ResponseWriter, dir http.File, p string) {
	files, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}
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

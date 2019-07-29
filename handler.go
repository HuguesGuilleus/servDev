// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main
import (
	"io"
	"log"
	"os"
	"net/http"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	// Get information
	url := r.URL.Path
	file,err1 := os.Open("."+url)
	if err1 != nil {
		log.Print("Error ",url)
		handle404(w,r)
		return
	}
	defer file.Close()
	info,err2 := file.Stat()
	if err2 != nil {
		log.Print("Error ",url)
		handle404(w,r)
		return
	}
	// Response
	if info.IsDir() {
		if url[len(url)-1] == '/' {
			log.Print("Directory ",url)
			handleDir(w,r,file)
		} else {
			log.Print("Redirect directory ",url)
			handleRedirect(w,url+"/")
		}
	} else {
		log.Print("File ",url)
		handleFile(w,file)
	}
}

func handleFile(w http.ResponseWriter, file *os.File) {
	w.Header().Add("Content-type",mime(file.Name()))
	io.Copy(w,file)
}

func handleDir(w http.ResponseWriter, r *http.Request, dir *os.File) {
	w.Write([]byte("Ceci est un répertoire, pas dev ;)\n"))
}

func handleRedirect(w http.ResponseWriter, url string) {
	w.Header().Add("Location",url)
	w.WriteHeader(301)
	w.Write([]byte(url+"\n"))
}

func handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("404 Not found /// Non trouvé /// No buscado\n"))
	w.Write([]byte(r.URL.String()+"\n"))
}

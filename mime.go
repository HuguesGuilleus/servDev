// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main
import (
	"log"
	"regexp"
)

var mimeExt *regexp.Regexp

func initMime() {
	mimeExt = regexp.MustCompile(".*[./](.*)")
}

// fonction qui done le type MIME en fonction du chemin/nom d'un fichier
func mime(path string) string {
	switch ext := mimeExt.ReplaceAllString(path,"$1"); ext {
		// Text
		case "html": return "text/html; charset=utf-8"
		case "txt": return "text/plain; charset=utf-8"
		case "css": return "text/css; charset=utf-8"
		case "webmanifest": return "application/manifest+json"
		case "vtt": return "text/vtt"
		case "pdf": return "application/pdf"
		case "md","markdown": return "text/markdown"
		case "tex": return "application/x-tex"
		// Media
		case "mp3": return "audio/mpeg"
		case "mp4": return "video/mp4"
		case "svg": return "image/svg+xml"
		case "ttf","tte","dfont": return "font/ttf"
		case "woff": return "application/font-woff"
		case "woff2": return "application/font-woff2"
		case "png": return "image/png"
		case "jpg","jpeg","JPG","JPEG": return "image/jpeg"
		// Divers
		case "js":		return "application/javascript"
		default:
			log.Print("MIME unknown: ",ext)
			return "text/plain; charset=utf-8"
	}
}

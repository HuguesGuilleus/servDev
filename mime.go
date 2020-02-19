// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"log"
	"regexp"
	"sync"
)

var (
	mimeExt   = regexp.MustCompile(".*[./](.*)")
	mimeMutex = sync.Mutex{}
	mimeList  = map[string]string{
		// Text
		"html":        "text/html; charset=utf-8",
		"txt":         "text/plain; charset=utf-8",
		"css":         "text/css; charset=utf-8",
		"webmanifest": "application/manifest+json",
		"vtt":         "text/vtt",
		"pdf":         "application/pdf",
		"md":          "text/markdown",
		"markdown":    "text/markdown",
		"tex":         "application/x-tex",
		// Media,
		"mp3":   "audio/mpeg",
		"mp4":   "video/mp4",
		"svg":   "image/svg+xml",
		"ttf":   "font/ttf",
		"tte":   "font/ttf",
		"dfont": "font/ttf",
		"woff":  "application/font-woff",
		"woff2": "application/font-woff2",
		"png":   "image/png",
		"jpg":   "image/jpeg",
		"jpeg":  "image/jpeg",
		"JPG":   "image/jpeg",
		"JPEG":  "image/jpeg",
		// Divers,
		"js":   "application/javascript",
		"mjs":  "application/javascript",
		"wasm": "application/wasm",
		"zip":  "application/zip",
		"json": "application/json; charset=utf-8",
		"bash": "text/plain; charset=utf-8",
		"sh":   "text/plain; charset=utf-8",
		"zsh":  "text/plain; charset=utf-8",
	}
)

// fonction qui done le type MIME en fonction du chemin/nom d'un fichier
func mime(path string) string {
	ext := mimeExt.ReplaceAllString(path, "$1")
	mimeMutex.Lock()
	defer mimeMutex.Unlock()
	m, exist := mimeList[ext]
	if !exist {
		log.Print("MIME unknown: ", ext)
		return "text/plain; charset=utf-8"
	}
	return m
}

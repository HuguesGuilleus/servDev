// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func init() {
	fmt.Print("\033[H\033[2J")
	log.SetFlags(log.Ltime)
}

func main() {
	addrHttp := flag.String("h", ":8000", `L'adresse d'écoute du server HTTP`)
	flag.Parse()
	root := flag.Arg(0)

	s := &server{Dir: http.Dir(root)}
	http.Handle("/", s)
	http.HandleFunc("/favicon.ico", s.favicon)

	r, _ := filepath.Abs(root)
	log.Printf("[LISTEN] %s on %s\n", r, *addrHttp)
	log.Fatal(http.ListenAndServe(*addrHttp, nil))
}

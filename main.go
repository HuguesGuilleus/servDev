// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func init() {
	fmt.Print("\033[H\033[2J")
	log.SetFlags(log.Ltime)
}

func main() {
	addrHttp := flag.String("http", ":8000", `L'adresse d'écoute du server HTTP`)
	root := flag.String("root", ".", `Le répertoire racine du server`)
	flag.Parse()

	s := &server{Dir: http.Dir(*root)}
	http.Handle("/", s)
	http.HandleFunc("/favicon.ico", s.favicon)

	log.Println("[LISTEN]", *addrHttp)
	log.Fatal(http.ListenAndServe(*addrHttp, nil))
}

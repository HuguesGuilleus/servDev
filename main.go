// HTTP Server for developping project
// 2019, 2020 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func init() {
	fmt.Print("\033[H\033[2J")
	log.SetFlags(log.Ltime)
}

func main() {
	addrHttp := flag.String("a", ":8000", `L'adresse d'écoute du server HTTP`)
	addrHttps := flag.String("s", ":8443", `L'addresse d'écoute HTTPS si il y a une clé et un certificat`)
	key := flag.String("k", "", "La clé de chiffrement (PEM)")
	cert := flag.String("c", "", "Le certificat")
	flag.Parse()
	root, _ := filepath.Abs(flag.Arg(0))

	s := &server{Dir: http.Dir(root)}
	http.Handle("/", s)
	http.HandleFunc("/favicon.ico", s.favicon)

	if *key != "" && *cert != "" {
		go func() {
			log.Printf("[LISTEN HTTP] %s (redirect)\n", *addrHttp)
			to := "https://" + *addrHttps
			if strings.HasPrefix(*addrHttps, ":") {
				to = "https://localhost" + *addrHttps
			}
			log.Fatal(http.ListenAndServe(
				*addrHttp,
				http.RedirectHandler(to, http.StatusTemporaryRedirect),
			))
		}()
		go func() {
			log.Printf("[LISTEN HTTPS] %s on %s\n", root, *addrHttps)
			log.Fatal(http.ListenAndServeTLS(*addrHttps, *cert, *key, nil))
		}()
		select {}
	} else {
		log.Printf("[LISTEN HTTP] %s on %s\n", root, *addrHttp)
		log.Fatal(http.ListenAndServe(*addrHttp, nil))
	}
}

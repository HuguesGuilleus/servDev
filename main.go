// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const VERSION = 1.5

func init() {
	fmt.Print("\033[H\033[2J")
	log.SetFlags(log.Ltime)
}

func main() {
	log.Println("[LISTEN]")
	s := &server{
		Dir: ".",
	}
	http.Handle("/", s)
	http.HandleFunc("/favicon.ico", s.favicon)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main_1() {
	fmt.Printf("\033[H\033[2J")

	log.Print("[INFO] HTTP Server for developping project // version: ", VERSION)
	// http.HandleFunc("/", handleMain)
	http.HandleFunc("/favicon.ico", handleFavicon)

	if len(os.Args) == 3 {
		// go func() {
		// 	err := (&http.Server{
		// 		Handler: redirect("https://localhost:8443"),
		// 		Addr:    ":8000",
		// 	}).ListenAndServe()
		// 	fatal(err)
		// }()
		// go fatal(http.ListenAndServe(":8000", redirect("https://localhost:8443")))

		log.Println("[INFO] Crt file:", os.Args[1])
		log.Println("[INFO] Key file:", os.Args[2])
		log.Println("")

		ca, err := ioutil.ReadFile("/home/hugues/Bureau/hugues.crt")
		fatal(err)

		crt := x509.NewCertPool()
		ok := crt.AppendCertsFromPEM(ca)
		log.Println("add cert:", ok)

		err = (&http.Server{
			Addr: ":8443",
			TLSConfig: &tls.Config{
				RootCAs: crt,
				// RootCAs: NewCertPool,
				ServerName: "localhost",
			},
		}).ListenAndServeTLS(os.Args[1], os.Args[2])
		// }).ListenAndServe()
		fatal(err)

		// fatal(http.ListenAndServeTLS(":8443", os.Args[1], os.Args[2], nil))
		// fatal(http.ListenAndServeTLS("https://localhost:8443/", os.Args[1], os.Args[2], nil))
	} else {
		fatal(http.ListenAndServe(":8000", nil))
	}
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

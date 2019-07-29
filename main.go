// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main
import (
	"fmt"
	"log"
	"net/http"
)

const VERSION =	1.1

func main() {
	fmt.Printf("\033[H\033[2J")
	initMime()
	log.Print("HTTP Server for developping project // version:",VERSION)
	http.HandleFunc("/",handleMain)
	err := http.ListenAndServe(":8000",nil)
	fmt.Println(err)
}

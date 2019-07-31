// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main
import (
	"fmt"
	"log"
	"net/http"
	"./templ"
	"os"
)

const VERSION =	1.4

func main() {
	fmt.Printf("\033[H\033[2J")
	initMime()

	if len(os.Args) > 1 && os.Args[1] == "-d" {
		l,err:=os.OpenFile("serv.log",os.O_WRONLY|os.O_APPEND|os.O_CREATE,0664)
		if err != nil {
			panic(err)
		}
		defer l.Close()
		log.SetOutput(l)
		log.Print("enregistrement du programme killServHttp.bash")
		saveKill()
	}

	log.Print("HTTP Server for developping project // version: ",VERSION)
	http.HandleFunc("/",handleMain)
	err := http.ListenAndServe(":8000",nil)
	fmt.Println(err)
}

func saveKill() {
	prog,err := os.OpenFile("killServHttp.bash",os.O_CREATE|os.O_WRONLY,0774)
	if err != nil {
		panic(err)
	}
	defer prog.Close()
	templ.Kill.Execute(prog,os.Getpid())
}

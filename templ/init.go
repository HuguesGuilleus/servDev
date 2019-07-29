// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package templ
import (
	"sync"
	"text/template"
)

var Redirect *template.Template
var Dir *template.Template
var NotFound *template.Template
var Kill *template.Template

func init() {
	wg := &sync.WaitGroup{}
	wg.Add(4)
	defer wg.Wait()

	go run(wg,&redirectSrc,&Redirect)
	go run(wg,&dirSrc,&Dir)
	go run(wg,&notFoundSrc,&NotFound)
	go run(wg,&killSrc,&Kill)
}

func run(wg *sync.WaitGroup, src *string, templ **template.Template) {
	*templ = template.Must(template.New("").Parse((*src)[1:]))
	*src = ""
	wg.Done()
}

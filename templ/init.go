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

func init() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	defer wg.Wait()

	go run(wg,&redirectSrc,&Redirect)
	go run(wg,&dirSrc,&Dir)
	go run(wg,&notFoundSrc,&NotFound)
}

func run(wg *sync.WaitGroup, src *string, templ **template.Template) {
	*templ = template.Must(template.New("").Parse((*src)[1:]))
	*src = ""
	wg.Done()
}

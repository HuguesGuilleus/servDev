// HTTP Server for developping project
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package templ

var redirectSrc string = `
FILE:redirect
`

var dirSrc string =`
FILE:dir
`

var notFoundSrc string = `
FILE:notFound
`

var killSrc string = `
#!/bin/bash

# HTTP Server for developping project
# 2019 GUILLEUS Hugues <ghugues@netc.fr>
# BSD 3-Clause "New" or "Revised" License

PID={{.}}

echo "Kill process number $PID"
kill -s KILL $PID

rm $0
`

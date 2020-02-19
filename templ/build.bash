#!/bin/bash

for templ in *.gohtml
do
	var=$(tr '[:lower:]' '[:upper:]' <<< ${templ:0:1})$( sed 's/.gohtml//' <<< ${templ:1})
	content=$( minify $templ --type html |
		sed 's/\\/\\\\/g' | sed 's/"/\\"/g' | tr '\n' ' ' )

	echo $var
	echo "
	// HTTP Server for developping project
	// 2019 GUILLEUS Hugues <ghugues@netc.fr>
	// BSD 3-Clause \"New\" or \"Revised\" License

	// Code generated for servHTTP DO NOT EDIT.

	package templ

	import (
		\"text/template\"
	)

	var $var = template.Must(template.New(\"\").Parse(\"$content\"))

	" > $( sed 's/gohtml/go/'<<< $templ)
done

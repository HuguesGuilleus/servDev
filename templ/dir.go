
	// HTTP Server for developping project
	// 2019 GUILLEUS Hugues <ghugues@netc.fr>
	// BSD 3-Clause "New" or "Revised" License

	// Code generated for servHTTP DO NOT EDIT.

	package templ

	import (
		"text/template"
	)

	var Dir = template.Must(template.New("").Parse("<!doctype html><meta charset=utf-8><meta name=viewport content=\"width=device-width,initial-scale=1\"><style>body{color:#000;background:#f9f9fa;font-family:monospace;font-size:x-large}.file{display:block;color:inherit;text-decoration:none}.file:hover{background:#1e90ff}.parent{margin-bottom:1ex}.fileName{display:inline-block;min-width:15ex;margin-right:1.5ex;text-align:right;font-weight:700}.fileSize::after{content:\"o\";display:inline-block;margin-left:.4ex;color:dimgray}address{margin-top:3ex;font-size:small}</style><a href=/ class=file><span class=fileName>/</span> [[ ROOT / RACINE ]]</a> <a href=../ class=\"file parent\"><span class=fileName>../</span> [[ PARENT / PADRES ]]</a> {{range .}} {{if .IsDir}} <a href={{.Name}}/ class=file><span class=fileName>{{html .Name}}/</span></a> {{else}} <a href={{.Name}} class=file><span class=fileName>{{html .Name}}</span> <span class=fileSize>{{html .Size}}</span></a> {{end}} {{else}}<p>EN: The are not file in this directory.<p>FR: Il n'y a pas de fichiers dans ce répertoire.</p>{{end}}<address>HTTP Server for developping project<br>BSD 3-Clause \"New\" or \"Revised\" License</address>"))

	

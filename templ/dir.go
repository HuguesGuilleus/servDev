
	// HTTP Server for developping project
	// 2019 GUILLEUS Hugues <ghugues@netc.fr>
	// BSD 3-Clause "New" or "Revised" License

	// Code generated for servHTTP DO NOT EDIT.

	package templ

	import (
		"text/template"
	)

	var Dir = template.Must(template.New("").Parse("<!doctype html><meta charset=utf-8><meta name=viewport content=\"width=device-width,initial-scale=1\"><style>body{color:#000;background:#f9f9fa;font-family:monospace;font-size:x-large}.fileList{display:grid}.file{display:inline-grid;grid-template-columns:10ex 1fr;grid-gap:1.5ex;color:inherit;text-decoration:none}.file:hover{background:#1e90ff}.fileSize{grid-column:1/2;text-align:right}.fileSize::after{content:\"o\";display:inline-block;margin-left:.4ex;color:dimgray}.fileName{grid-column:2/3}.parent{margin-bottom:1ex}address{margin-top:3ex;font-size:small}</style><main class=fileList><a href=/ class=file><span class=fileName>/</span></a> <a href=../ class=\"file parent\"><span class=fileName>..</span></a> {{- range . -}} {{- if .IsDir -}} <a href={{.Name}}/ class=file><span class=fileName>{{html .Name}}/</span></a> {{- else -}} <a href={{.Name}} class=file><span class=fileSize>{{html .Size}}</span> <span class=fileName>{{html .Name}}</span></a> {{- end -}} {{- else -}}<p>EN: The are not file in this directory.<p>FR: Il n'y a pas de fichiers dans ce répertoire.</p>{{- end -}}</main><address>HTTP Server for developping project<br>BSD 3-Clause \"New\" or \"Revised\" License</address>"))

	

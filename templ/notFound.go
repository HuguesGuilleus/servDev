
	// HTTP Server for developping project
	// 2019 GUILLEUS Hugues <ghugues@netc.fr>
	// BSD 3-Clause "New" or "Revised" License

	// Code generated for servHTTP DO NOT EDIT.

	package templ

	import (
		"text/template"
	)

	var NotFound = template.Must(template.New("").Parse("<!doctype html><meta charset=utf-8><meta name=viewport content=\"width=device-width,initial-scale=1\"><style>body{background:#f9f9fa;font-family:monospace;font-size:large}a{font-size:xx-large}a:hover{background:#1e90ff}address{margin-top:3ex;font-size:small}</style><h1>404</h1><p>Not found /// Non trouvé /// No buscado</p>{{ $last := 0 }} {{range $i,$e := . }} {{ $last = $i }} {{end}} {{ $url := \"\" }} {{range $i,$e := .}} {{if ne $i $last }} {{ $url = print $url $e \"/\" }} <a href={{$url}}>{{html $e}}/</a> {{else if ne $e \"\"}} {{ $url = print $url $e }} <a href={{$url}}>{{html $e}}</a> {{end}} {{end}}<address>HTTP Server for developping project<br>BSD 3-Clause \"New\" or \"Revised\" License</address>"))

	

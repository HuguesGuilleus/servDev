// Code generated by staticFile; DO NOT EDIT.
// source files: ["templ/dir.gohtml" "templ/notFound.gohtml"]
// dev mode: false

package main

func TemplDir() []byte { return []byte("<!doctype html><html><head><meta charset=utf-8><meta name=viewport content=\"width=device-width,initial-scale=1\"><style>*[hidden]{display:none!important}body{color:#000;background:#f9f9fa;font-family:monospace;font-size:x-large}.fileList{display:grid}.file{display:inline-grid;grid-template-columns:12ex 1fr;grid-gap:1.5ex;color:inherit;text-decoration:none}.file:hover{background:#1e90ff}.fileSize{grid-column:1/2;text-align:right}.fileSize::after{content:\"o\";display:inline-block;margin-left:.4ex;color:dimgray}.fileName{grid-column:2/3;font-weight:bolder}.parent{margin-bottom:1ex}#search{display:block;width:100%;margin-bottom:3ex}address{margin-top:3ex;font-size:small}</style></head><body><input type=search id=search autofocus><main class=fileList><a href=/ class=file><div class=fileName>/</div></a><a href=../ class=\"file parent\"><div class=fileName>..</div></a>{{- range . -}}\n{{- if .IsDir -}}\n<a href=\"{{urlquery .Name}}/\" class=file><span class=fileName>{{html .Name}}/</span></a>\n{{- else -}}\n<a href=\"{{urlquery .Name}}\" class=file><div class=fileSize>{{html .Size}}</div><div class=fileName>{{html .Name}}</div></a>{{- end -}}\n{{- else -}}<p>EN: The are not file in this directory.</p><p>FR&nbsp;: Il n'y a pas de fichiers dans ce répertoire.</p>{{- end -}}</main><address>HTTP Server for developping project<br>BSD 3-Clause \"New\" or \"Revised\" License</address><script>(()=>{let s=document.getElementById(\"search\");s.addEventListener(\"input\",()=>document.querySelectorAll(\".file\").forEach(file=>file.hidden=!file.querySelector(\".fileName\").textContent.includes(s.value)));})();</script></body></html>") }

func TemplNotFound() []byte { return []byte("<!doctype html><html><head><meta charset=utf-8><meta name=viewport content=\"width=device-width,initial-scale=1\"><style>body{background:#f9f9fa;font-family:monospace;font-size:large}a{font-size:xx-large}a:hover{background:#1e90ff}address{margin-top:3ex;font-size:small}</style></head><body><h1>404</h1><p>Not found /// Non trouvé /// No buscado</p>{{- $last := 0 -}}\n{{- range $i,$e := . -}}\n{{- $last = $i -}}\n{{- end -}}\n{{- $url := \"\" -}}\n{{- range $i,$e := . -}}\n{{- if ne $i $last -}}\n{{- $url = print $url $e \"/\" -}}\n<a href=\"{{print $url}}\">{{html $e}}/</a>\n{{- else if ne $e \"\" -}}\n{{- $url = print $url $e -}}\n<a href=\"{{print $url}}\">{{html $e}}/</a>\n{{- end -}}\n{{- end -}}<address>HTTP Server for developping project<br>BSD 3-Clause \"New\" or \"Revised\" License</address></body></html>") }

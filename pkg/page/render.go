package page

import (
	"bytes"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("template.html"))

func renderTemplate(post Post) string {
	var buffer bytes.Buffer

	err := tmpl.Execute(&buffer, post)

	if err != nil {
		panic(err)
	}

	return string(buffer.Bytes())
}

func RenderPost(post Post) string {
	html := renderTemplate(post)

	return html
}

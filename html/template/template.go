package main

import (
	"os"
	"html/template"
)

func main() {
	tmpl, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
//	err = tmpl.Execute(os.Stdout, "<script>alert('you have been pwned')</script>")
}
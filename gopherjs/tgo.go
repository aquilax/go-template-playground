package main

//go:generate gopherjs build --minify

import (
	"bytes"
	"encoding/json"
	"text/template"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Get("window").Set("renderTemplate", render)
}

func render(t, data string) (string, string) {
	var v interface{}
	var err error
	var tmpl *template.Template
	err = json.Unmarshal([]byte(data), &v)
	if err != nil {
		return "", err.Error()
	}
	tmpl, err = template.New("test").Parse(t)
	if err != nil {
		return "", err.Error()
	}
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, v)
	if err != nil {
		return "", err.Error()
	}
	return buffer.String(), ""
}

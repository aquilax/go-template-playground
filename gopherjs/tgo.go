package main

//go:generate gopherjs build --minify

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"text/template"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Get("window").Set("renderTemplate", render)
	js.Global.Get("window").Set("compress", compress)
	js.Global.Get("window").Set("decompress", decompress)
}

// renders template
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

func compress(text string) (string, string) {
	var b bytes.Buffer
	w, err := flate.NewWriter(&b, flate.BestCompression)
	if err != nil {
		return "", err.Error()
	}
	w.Write([]byte(text))
	w.Close()

	b.Bytes()
	encoded := base64.RawURLEncoding.EncodeToString(b.Bytes())
	return encoded, ""
}

func decompress(text string) (string, string) {
	b, err := base64.RawURLEncoding.DecodeString(text)
	if err != nil {
		return "", err.Error()
	}
	r := flate.NewReader(bytes.NewReader(b))
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err.Error()
	}
	return string(data), ""
}

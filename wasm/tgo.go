package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/json"
	"io"
	"text/template"
)

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

	encoded := base64.RawURLEncoding.EncodeToString(b.Bytes())
	return encoded, ""
}

func decompress(text string) (string, string) {
	b, err := base64.RawURLEncoding.DecodeString(text)
	if err != nil {
		return "", err.Error()
	}
	r := flate.NewReader(bytes.NewReader(b))
	data, err := io.ReadAll(r)
	if err != nil {
		return "", err.Error()
	}
	return string(data), ""
}

package utils

import (
	"bytes"
	"html/template"
)

func TemplateParse(name, text string, data interface{}) ([]byte, error) {
	t, err := template.New(name).Parse(text)
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	if err := t.Execute(&out, data); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func TemplateParseFile(filename string, data interface{}) ([]byte, error) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	if err := t.Execute(&out, data); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

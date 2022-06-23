package main

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

type Model struct {
	Variables map[string]any
}

func renderTemplate(templateFile string, data *Model) ([]byte, error) {
	originalContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return nil, err
	}

	tmpl := template.New(templateFile)

	t, err := tmpl.Parse(string(originalContent))
	if err != nil {
		return nil, err
	}

	renderedContent := new(bytes.Buffer)
	err = t.Execute(renderedContent, data.Variables)
	if err != nil {
		return nil, err
	}

	return renderedContent.Bytes(), nil
}

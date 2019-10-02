package template

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"
)

// NewTemplate returns a new template instance
func NewTemplate() *Template {
	return &Template{}
}

// Template is responsible for producing files from the given files
type Template struct{}

// ExecuteFile executes the given template file and returns the final result
func (t *Template) ExecuteFile(file string) (string, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return t.Execute(string(f))
}

// Execute executes the given template and returns the final result
func (t *Template) Execute(in string) (string, error) {
	funcMap := template.FuncMap{
		"env": os.Getenv,
	}

	tmpl, err := template.New("file").
		Funcs(funcMap).
		Parse(string(in))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

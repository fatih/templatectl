package template

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

// NewTemplate returns a new template instance
func NewTemplate() *Template {
	return &Template{LookupFunc: os.LookupEnv}
}

// Template is responsible for producing files from the given files
type Template struct {
	LookupFunc func(string) (string, bool)
}

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
		"env": func(key string) (string, error) {
			val, found := t.LookupFunc(key)
			if !found {
				return "", fmt.Errorf("environment variable %q is not defined", key)
			}
			return val, nil
		},
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

package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/templatectl/template"
)

func main() {
	if err := realMain(); err != nil {
		log.Fatalln(err)
	}
}

func realMain() error {
	input := flag.String("input", "", "File path to process")
	output := flag.String("output", "", "Path to save the processed template file(optional)")
	flag.Parse()

	if *input == "" {
		return errors.New("usage: templatectl --input file.tmpl --output exported.txt")
	}

	t := template.NewTemplate()
	buf, err := t.ExecuteFile(*input)
	if err != nil {
		return fmt.Errorf("error executing file %q template: %s", *input, err)
	}

	if *output != "" {
		info, err := os.Stat(*input)
		if err != nil {
			return fmt.Errorf("error retrieving file info %q: %s", *output, err)
		}

		err = ioutil.WriteFile(*output, []byte(buf), info.Mode())
		if err != nil {
			return fmt.Errorf("error saving processed file %q: %s", *output, err)
		}
	} else {
		fmt.Print(buf)
	}

	return nil
}

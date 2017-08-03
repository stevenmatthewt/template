package template

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func (t Templater) templateFile(path string) error {
	destPath := filepath.Join(t.destination, path)
	contents, err := t.readFile(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %s", err)
	}
	if strings.HasSuffix(path, ".tmpl") {
		destPath = destPath[:len(destPath)-5]
		contents, err = parseTemplate(path, t.data)
		if err != nil {
			return fmt.Errorf("failed to apply template to file: %s", err)
		}
	}

	file, err := os.Create(destPath)
	defer file.Close()
	_, err = file.Write(contents)
	if err != nil {
		return fmt.Errorf("failed to write to file: %s", err)
	}

	return nil
}

func (t Templater) readFile(path string) (contents []byte, err error) {
	contents, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func parseTemplate(templatePath string, data interface{}) (result []byte, err error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %s", err)
	}

	buffer := &bytes.Buffer{}
	err = t.Execute(buffer, data)
	return buffer.Bytes(), err
}

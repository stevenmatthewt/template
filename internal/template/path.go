package template

import (
	"bytes"
	"html/template"
	"path/filepath"
	"strings"
)

func (t Templater) finalPath(path string) (string, error) {
	var parsedList []string
	for path != "" && path != "." && path != "/" {
		current := filepath.Base(path)
		path = filepath.Dir(path)
		parsed, err := parseTemplateString(current, t.data)
		if err != nil {
			return "", err
		}
		if parsed == "" {
			return "", nil
		}
		parsedList = append([]string{parsed}, parsedList...)
	}
	return filepath.Join(parsedList...), nil
}

func parseTemplateString(str string, data interface{}) (result string, err error) {
	if !strings.Contains(str, "{{") {
		return str, nil
	}
	t := template.New(str)
	t.Parse(str)

	buffer := &bytes.Buffer{}
	err = t.Execute(buffer, data)
	return string(buffer.Bytes()), err
}

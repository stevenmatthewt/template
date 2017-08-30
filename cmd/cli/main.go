package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/segmentio/go-prompt"
	"github.com/stevenmatthewt/template/internal/template"
	"gopkg.in/yaml.v2"
)

type CLIFlags struct {
	TemplatePath    string
	DestinationPath string
}

type Config struct {
	Prompts map[string]Prompt `yaml:"prompts"`
}

type Prompt struct {
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
	Type        string `yaml:"type"`
}

func main() {
	flags := getFlags()

	config, err := getConfig(flags.TemplatePath)
	if err != nil {
		panic(err)
	}
	templateData, err := getTemplateData(config)
	if err != nil {
		panic(err)
	}
	fn := template.New(flags.TemplatePath, flags.DestinationPath, templateData)
	err = fn.Walk(flags.TemplatePath, fn)
	if err != nil {
		panic(err)
	}
}

func getFlags() (flags CLIFlags) {
	flag.StringVar(&flags.TemplatePath, "template", "", "Location of the template to use.")
	flag.StringVar(&flags.DestinationPath, "destination", "./", "Destination folder of new project.")
	flag.Parse()

	return flags
}

func getConfig(path string) (config Config, err error) {
	bytes, err := ioutil.ReadFile(filepath.Join(path, "tmpl.yml"))
	if err != nil {
		return Config{}, fmt.Errorf("failed to read template config file: %s", err)
	}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to parse template config file: %s", err)
	}

	return config, nil
}

func getTemplateData(config Config) (map[string]interface{}, error) {
	var err error
	data := make(map[string]interface{})
	for name, p := range config.Prompts {
		input := prompt.String(p.Description)
		switch p.Type {
		case "bool":
			boolMap := map[string]bool{
				"true":  true,
				"True":  true,
				"false": false,
				"False": false,
				"Y":     true,
				"N":     false,
				"y":     true,
				"n":     false,
				"Yes":   true,
				"No":    false,
			}
			var ok bool
			data[name], ok = boolMap[input]
			if !ok {
				return nil, fmt.Errorf("input is not recognized as a boolean: %s", input)
			}
		case "string":
			data[name] = input
		case "":
			data[name] = input
		case "int":
			data[name], err = strconv.Atoi(input)
			if err != nil {
				return nil, fmt.Errorf("input is not recognized as an integer: %s", input)
			}
		}
	}

	return data, nil
}

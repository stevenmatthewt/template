package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/segmentio/go-prompt"
	"github.com/stevenmatthewt/template/internal/template"
	"github.com/stevenmatthewt/template/internal/walk"
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
}

func main() {
	flags := getFlags()

	config, err := getConfig(flags.TemplatePath)
	if err != nil {
		panic(err)
	}
	templateData := getTemplateData(config)
	fn := template.New(flags.DestinationPath, templateData)
	walk.Walk(flags.TemplatePath, fn)
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

func getTemplateData(config Config) map[string]interface{} {
	data := make(map[string]interface{})
	for name, p := range config.Prompts {
		data[name] = prompt.String(p.Description)
	}

	return data
}

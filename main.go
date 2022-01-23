package main

import (
	"blog-poster/app"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	config := readConfig()
	blog := app.NewTistory(config.BlogName, config.AccessToken)
	blog.List()
}

func readConfig() Config {
	filename, _ := filepath.Abs("config.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	return config
}

type Config struct {
	AccessToken string `yaml:"AccessToken"`
	BlogName    string `yaml:"BlogName"`
}

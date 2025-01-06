package framework

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "application.yaml"

type Cfg struct {
	PostgresUrl struct {
		URI string `yaml:"uri"`
	} `yaml:"postgres"`
	ServerPort struct {
		PORT string `yaml:"port"`
	} `yaml:"server"`
}

var AppConfig *Cfg

func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)

	if err != nil {
		fmt.Println(err)
	}
}

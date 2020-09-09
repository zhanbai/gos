package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Address string `yaml:"address"`
	}
	Database struct {
		Link string `yaml:"link"`
	}
}

func Get() Config {
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

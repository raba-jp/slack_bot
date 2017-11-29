package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	BotID    string   `yaml:"bot_id"`
	Channels []string `yaml:"channels"`
}

func NewConfig() (Config, error) {
	buf, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		return nil, err
	}
	var config Config
	yaml.Unmarshal(buf, &config)
	return config, nil
}

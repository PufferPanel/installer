package main

import (
	"encoding/json"
	"github.com/pufferpanel/installer/resources"
	"io/ioutil"
)

var config *resources.Config
var loaded = false

func GetConfig() (resources.Config, error) {
	if loaded {
		return *config, nil
	}
	err := LoadConfig()
	return *config, err
}

func LoadConfig() error {
	config = &resources.Config{}

	loaded = true

	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, config)
	return err
}

package resources

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Database MysqlConfig `json:"mysql"`
}

type MysqlConfig struct {
	Host     string `json:"host"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
}

var config *Config

func Get() (Config, error) {
	if config == nil {
		config = &Config{
			Database: MysqlConfig{
				Host:     "localhost",
				Database: "pufferpanel",
				Port:     "3306",
			},
		}

		file, err := ioutil.ReadFile("config.json")
		if err != nil {
			return *config, err
		}

		err = json.Unmarshal(file, config)
		return *config, err
	}

	return *config, nil
}

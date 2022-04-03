package config

import (
	yaml "gopkg.in/yaml.v2"
	"os"
)

var _config *config

type config struct {
	Debug bool   `yaml:"debug"`
	MySQL string `yaml:"mysql"`
}

func Debug() bool {
	return _config.Debug
}

func MySQL() string {
	return _config.MySQL
}

func init() {
	_config = &config{}

	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic("Cannot load config file: " + err.Error())
	}

	err = yaml.Unmarshal(configFile, _config)
	if err != nil {
		panic("Cannot load config file: " + err.Error())
	}
}

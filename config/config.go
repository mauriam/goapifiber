package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Port string `yaml:"PORT"`
}

// GetConfig carga las configuraciones del archivo Config.yaml
// retorna las configuraciones contenidas en el archivo y el err si existe
func GetConfig() (conf *Settings, err error) {
	data, err := ioutil.ReadFile("../config/config.yaml")
	if err != nil {
		return
	}
	conf = &Settings{}
	err = yaml.Unmarshal(data, conf)
	return
}

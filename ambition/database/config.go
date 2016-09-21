package database

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	DBName       string
	DBUser       string
	DBPassword   string
	DBLocal      bool
	DBHost       string
	DBPort       int
	DBSSL        string
	AmbitionPort int
}

func ReadConfiguration(file string) Configuration {

	configJson, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var config Configuration
	err = json.Unmarshal(configJson, &config)

	return config

}

package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/meson10/cdnlysis_engine/client"
	"gopkg.in/yaml.v2"
)

type InfluxConfig struct {
	initalized bool

	Influx client.ClientConfig
}

const baseConfig = `
influx:
    url: 127.0.0.1:8086
    username: cdnlysis
    password: cdnlysis
    isudp: true
    database: cdnlogs
`

func cliArgs() string {
	confFlag := flag.Lookup("config")
	if confFlag == nil || len(confFlag.Value.String()) == 0 {
		return ""
	}

	return confFlag.Value.String()
}

var Settings InfluxConfig

func GetConfig() InfluxConfig {
	if Settings.initalized {
		return Settings
	}

	//Try to look for a module level cached value
	var path string = cliArgs()

	conf := InfluxConfig{}
	err := yaml.Unmarshal([]byte(baseConfig), &conf)
	if err != nil {
		log.Println(baseConfig)
		log.Println("1", err)
	}

	if path != "" {
		confData := readYaml(path)
		err := yaml.Unmarshal(confData, &conf)
		if err != nil {
			log.Println("2", err)
		}
	}

	//Set the module level cached value.
	Settings = conf
	return Settings
}

func readYaml(path string) []byte {
	//Load YAML file from the path provided.
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return yamlFile
}

package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/cdnlysis/cdnlysis.v1/conf"

	"github.com/cdnlysis/cdnlysis_influxdb/client"
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

var Settings InfluxConfig

func GetConfig() InfluxConfig {
	if Settings.initalized {
		return Settings
	}

	//Try to look for a module level cached value
	var path string = conf.CliArgs()

	conf := InfluxConfig{}
	err := yaml.Unmarshal([]byte(baseConfig), &conf)
	if err != nil {
		log.Println("1", err)
	}

	if path != "" {
		confData := readYaml(path)
		err := yaml.Unmarshal(confData, &conf)
		if err != nil {
			log.Println("2", err)
		}
	}

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

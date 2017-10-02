// Created by nazarigonzalez on 2/10/17.

package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

var configFile = flag.String("config", "development.config.toml", "Config file")

var Data = func() *configObject {
	flag.Parse()

	if _, err := os.Stat(*configFile); os.IsNotExist(err) {
		if err := ioutil.WriteFile(*configFile, []byte(example), 0644); err != nil {
			panic(err)
		}
	}

	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}

	c := configObject{}
	_, err = toml.Decode(string(data), &c)
	if err != nil {
		panic(err)
	}

	return &c
}()

type configObject struct {
	Name     string
	Port     int
	Database database
	Logger   logger
}

type logger struct {
	Level     int
	File      bool
	Path      string
	FileLevel int `toml:"file_level"`
}

type database struct {
	Host     string
	User     string
	Name     string
	SSLMode  string
	Password string
	Debug    bool
}

var example = `
name = "Tarentola"      #app name
port = 8000             #http port

[database]
host = "localhost"
user = "postgres"
name = "tarentola"
sslmode = "disable"
password = "postgres"
debug = false           #log database actions

[logger]
level = 0               #[trace, debug, info, log, warn, error, fatal]
path = "./logs"
file = true
file_level = 0          #same as level but for file output
`

// Created by nazarigonzalez on 2/10/17.

package config

import "github.com/BurntSushi/toml"

//todo read os.Flags
//todo load toml

var Data = func() *configObject {
	c := configObject{}
	_, err := toml.Decode(example, &c)
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

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
	Name        string
	Port        int
	Environment string
	Static      static
	Database    database
	Logger      logger
	Auth        auth
	CORS        cors
}

func (c *configObject) IsProd() bool {
	return c.Environment == "production"
}

type logger struct {
	Async     bool
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

type auth struct {
	TokenExpire int `toml:"token_expire"`
	Secret      string
}

type cors struct {
	Enabled bool
	Origins []string
}

type static struct {
	Enabled bool
	Port    int
	Path    string
}

var example = `
name = "Tarentola"      				#app name
port = 8000             				#http port
environment = "development" 		#[production, others...]

[static]								#static files
enabled = true
port = 8050
path = "./static"

[database]
host = "localhost"
user = "postgres"
name = "tarentola"
sslmode = "disable"
password = "postgres"
debug = false           #log database actions

[logger]
async = false
level = 0               #[trace, debug, info, log, warn, error, fatal]
file = true
file_level = 0          #same as level but for file output
path = "./logs"

[auth]
token_expire = 3600			#seconds
secret = "please change me in production!"	#secret key to sign the auth token

[cors]									#cross domain requests
enabled = true
origins = ["*"]
`

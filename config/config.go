package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	DB struct {
		Database string
		URL      string
	} `toml:"db"`

	Info struct {
		Port    string
		Service string
		Log     string
	} `toml:"info"`

	Aws struct {
		IAMAKey   string
		IAMSKey   string
		Region    string
		Bucket    string
		S3BaseURL string
	} `toml:"aws"`

	Redis struct {
		Url      string
		DB       int
		User     string
		Password string
	} `toml:"redis"`
}

func NewConfig(path string) *Config {
	c := new(Config)

	if f, err := os.Open(path); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}

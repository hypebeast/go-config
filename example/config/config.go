package config

import (
	"github.com/hypebeast/go-config/config"
)

type BaseConfig struct {
	Host    string
	Port    int
	ApiKey  string
	Twitter struct {
		Username string
		Password string
	}
	Facebook struct {
		Username string
		Password string
	}
}

type MongoConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	OptionInBase string
}

var baseConfig BaseConfig
var mongoConfig MongoConfig

func init() {
	// Initialize the config system
	config.Init("config", "GOENV")

	// Get the base config
	config.Get("base", &baseConfig)
	// Get the MongoDB config
	config.Get("mongo", &mongoConfig)
}

func BaseConf() *BaseConfig {
	return &baseConfig
}

func MongoConf() *MongoConfig {
	return &mongoConfig
}

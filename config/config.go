// Package config holds the config needed for the service to run
package config

import (
	"github.com/alexflint/go-arg"
)

// Config is the configuration to run the service
// args are parsed from go-arg, https://github.com/alexflint/go-arg
type Config struct {
	Env  string `arg:"env:ENVIRONMENT" default:"production" validate:"required,oneof=test local sandbox dev staging production"`
	Port int    `arg:"env:PORT" default:"8080" validate:"required,lte=65536"`
}

// NewConfig return a new instance of Config
func NewConfig() Config {
	config := Config{}
	arg.MustParse(&config)

	return config
}

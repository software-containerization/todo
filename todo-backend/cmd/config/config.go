package config

import (
	"fmt"

	"github.com/google/wire"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	DSN  string
	Port string
}

var Set = wire.NewSet(
	ProvideConfig,
)

var c Configuration

// ProvideConfig provides the config object
func ProvideConfig() *Configuration {
	return &c
}

// Load populates the config object with env variables
func (c *Configuration) Load() error {
	c.DSN = viper.GetString("dsn")
	c.Port = viper.GetString("port")
	if c.Port == "" {
		log.Info("No env TODO_PORT found, using port 8080")
		c.Port = "8080"
	}

	if c.DSN == "" {
		return fmt.Errorf("No TODO_DSN environment variable found")
	}

	return nil
}

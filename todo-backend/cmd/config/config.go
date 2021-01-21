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
		err := c.LoadSeparateDBCredentials()
		if err != nil {
			return fmt.Errorf("No database credentials found in environment. Please set TODO_DSN or separate credentials.")
		}
	}

	log.Infof("Using DSN: %s", c.DSN)

	return nil
}

func (c *Configuration) LoadSeparateDBCredentials() error {
	dbName := viper.GetString("db_name")
	dbUser := viper.GetString("db_user")
	dbPasswd := viper.GetString("db_passwd")
	dbHost := viper.GetString("db_host")
	dbPort := viper.GetString("db_port")

	if dbName == "" {
		return fmt.Errorf("No TODO_DB_NAME environment variable found")
	}

	if dbUser == "" {
		return fmt.Errorf("No TODO_DB_USER environment variable found")
	}

	if dbPasswd == "" {
		return fmt.Errorf("No TODO_DB_PASSWD environment variable found")

	}

	if dbPort == "" {
		return fmt.Errorf("No TODO_DB_PORT environment variable found")
	}

	if dbHost == "" {
		return fmt.Errorf("No TODO_DB_HOSTenvironment variable found")
	}

	c.DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbHost, dbPort, dbName)

	return nil
}

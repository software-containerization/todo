//+build wireinject

package cmd

import (
	"todo-api/cmd/config"
	"todo-api/pkg/server"

	"github.com/google/wire"
)

func initialiseServer(c *config.Configuration) *server.Server {
	wire.Build(server.Set)
	return &server.Server{}
}

func initialiseConfig() *config.Configuration {
	wire.Build(config.Set)
	return nil
}

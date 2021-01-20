// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cmd

import (
	"github.com/gorilla/mux"
	"todo-api/cmd/config"
	"todo-api/pkg/server"
	"todo-api/pkg/storage"
)

// Injectors from wire.go:

func initialiseServer(c *config.Configuration) *server.Server {
	router := mux.NewRouter()
	db := storage.NewGormDB(c)
	itemRepository := storage.NewItemRepository(db)
	serverServer := server.NewServer(router, itemRepository)
	return serverServer
}

func initialiseConfig() *config.Configuration {
	configuration := config.ProvideConfig()
	return configuration
}
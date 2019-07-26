// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"github.com/{{github}}/internal/router"
	"github.com/{{github}}/internal/server"
	"github.com/{{github}}/internal/store/database"
	"github.com/{{github}}/internal/store/memory"
	"github.com/{{github}}/types"
)

// Injectors from wire.go:

func initServer(config *types.Config) (*server.Server, error) {
	db, err := database.ProvideDatabase(config)
	if err != nil {
		return nil, err
	}
	{{toLower child}}Store := database.Provide{{child}}Store(db)
	{{toLower parent}}Store := database.Provide{{parent}}Store(db)
	memberStore := database.ProvideMemberStore(db)
	projectStore := database.ProvideProjectStore(db)
	userStore := database.ProvideUserStore(db)
	systemStore := memory.New(config)
	handler := router.New({{toLower child}}Store, {{toLower parent}}Store, memberStore, projectStore, userStore, systemStore)
	serverServer := server.ProvideServer(config, handler)
	return serverServer, nil
}

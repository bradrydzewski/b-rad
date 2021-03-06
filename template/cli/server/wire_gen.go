// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"github.com/{{toLower repo}}/internal/cron"
	"github.com/{{toLower repo}}/internal/router"
	"github.com/{{toLower repo}}/internal/server"
	"github.com/{{toLower repo}}/internal/store/database"
	"github.com/{{toLower repo}}/internal/store/memory"
	"github.com/{{toLower repo}}/types"
)

// Injectors from wire.go:

func initSystem(config *types.Config) (*system, error) {
	db, err := database.ProvideDatabase(config)
	if err != nil {
		return nil, err
	}
	{{toLower child}}Store := database.Provide{{title child}}Store(db)
	{{toLower parent}}Store := database.Provide{{title parent}}Store(db)
	memberStore := database.ProvideMemberStore(db)
	{{toLower project}}Store := database.Provide{{title project}}Store(db)
	userStore := database.ProvideUserStore(db)
	systemStore := memory.New(config)
	handler := router.New({{toLower child}}Store, {{toLower parent}}Store, memberStore, {{toLower project}}Store, userStore, systemStore)
	serverServer := server.ProvideServer(config, handler)
	nightly := cron.NewNightly()
	serverSystem := newSystem(serverServer, nightly)
	return serverSystem, nil
}

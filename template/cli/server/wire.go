// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

//+build wireinject

package server

import (
	"github.com/{{toLower repo}}/internal/cron"
	"github.com/{{toLower repo}}/internal/router"
	"github.com/{{toLower repo}}/internal/server"
	"github.com/{{toLower repo}}/internal/store/database"
	"github.com/{{toLower repo}}/internal/store/memory"
	"github.com/{{toLower repo}}/types"

	"github.com/google/wire"
)

func initSystem(config *types.Config) (*system, error) {
	wire.Build(
		database.WireSet,
		memory.WireSet,
		router.WireSet,
		server.WireSet,
		cron.WireSet,
		newSystem,
	)
	return &system{}, nil
}

// Copyright 2019 Brad Rydzewski. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//+build wireinject

package server

import (
	"github.com/{{github}}/internal/router"
	"github.com/{{github}}/internal/server"
	"github.com/{{github}}/internal/store/database"
	"github.com/{{github}}/internal/store/memory"
	"github.com/{{github}}/types"

	"github.com/google/wire"
)

func initServer(config *types.Config) (*server.Server, error) {
	wire.Build(
		database.WireSet,
		memory.WireSet,
		router.WireSet,
		server.WireSet,
	)
	return &server.Server{}, nil
}

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

package server

import (
	"net/http"

	"github.com/{{github}}/types"

	"github.com/google/wire"
)

// WireSet provides a wire set for this package
var WireSet = wire.NewSet(ProvideServer)

// ProvideServer provides a server instance
func ProvideServer(config *types.Config, handler http.Handler) *Server {
	return &Server{
		Acme:    config.Server.Acme.Enabled,
		Addr:    config.Server.Bind,
		Host:    config.Server.Host,
		Handler: handler,
	}
}

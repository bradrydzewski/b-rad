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

// Package swagger defines the swagger specification.
//
//     Schemes: http, https
//     BasePath: /api/v1
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package swagger

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:generate swagger generate spec -o files/swagger.json

//go:embed files/*
var content embed.FS

// FileSystem provides access to the static web server
// content, embedded in the binary.
func FileSystem() http.FileSystem {
	fsys, _ := fs.Sub(content, "files")
	return http.FS(fsys)
}

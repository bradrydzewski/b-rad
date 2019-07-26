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

package user

import (
	"net/http"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/api/request"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"
)

// HandleProjects returns an http.HandlerFunc that writes a json-encoded
// list of projects to the response body.
func HandleProjects(projects store.ProjectStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())
		list, err := projects.List(r.Context(), viewer.ID, types.Params{})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("email", viewer.Email).
				Errorf("cannot list projects")
		} else {
			render.JSON(w, list, 200)
		}
	}
}

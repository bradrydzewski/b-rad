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

package {{toLower child}}s

import (
	"net/http"
	"strconv"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"

	"github.com/go-chi/chi"
)

// HandleFind returns an http.HandlerFunc that writes the
// json-encoded {{toLower parent}} details to the response body.
func HandleFind({{toLower parent}}s store.{{parent}}Store, {{toLower child}}s store.{{child}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectID, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		{{toLower parent}}ID, err := strconv.ParseInt(chi.URLParam(r, "{{toLower parent}}"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse {{toLower parent}} id")
			return
		}

		{{toLower child}}ID, err := strconv.ParseInt(chi.URLParam(r, "{{toLower child}}"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse {{toLower child}} id")
			return
		}

		{{toLower parent}}, err := {{toLower parent}}s.Find(r.Context(), {{toLower parent}}ID)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", {{toLower parent}}ID).
				Debugln("{{toLower parent}} not found")
			return
		}

		{{toLower child}}, err := {{toLower child}}s.Find(r.Context(), {{toLower child}}ID)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", {{toLower child}}ID).
				Debugln("{{toLower parent}} not found")
			return
		}

		if {{toLower parent}}.Project != projectID {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithField("{{toLower parent}}", {{toLower parent}}ID).
				WithField("{{toLower child}}", {{toLower child}}ID).
				WithField("project", projectID).
				Debugln("project id mismatch")
			return
		}

		if {{toLower parent}}.ID != {{toLower child}}.{{parent}} {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithField("{{toLower parent}}.id", {{toLower parent}}.ID).
				WithField("{{toLower child}}.id", {{toLower child}}.ID).
				WithField("project", projectID).
				Debugln("{{toLower parent}} id mismatch")
			return
		}

		render.JSON(w, {{toLower child}}, 200)
	}
}

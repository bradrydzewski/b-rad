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

package {{toLower parent}}s

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"

	"github.com/go-chi/chi"
)

// HandleUpdate returns an http.HandlerFunc that processes http
// requests to update the object details.
func HandleUpdate({{toLower parent}}s store.{{parent}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		project, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		id, err := strconv.ParseInt(chi.URLParam(r, "{{toLower parent}}"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse {{toLower parent}} id")
			return
		}

		in := new(types.{{parent}}Input)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("project", project).
				WithField("id", id).
				Debugln("cannot unmarshal json request")
			return
		}

		{{toLower parent}}, err := {{toLower parent}}s.Find(r.Context(), id)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", id).
				Debugln("{{toLower parent}} not found")
			return
		}

		if {{toLower parent}}.Project != project {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithField("id", id).
				WithField("project", project).
				Debugln("project id mismatch")
			return
		}

		if in.Name.IsZero() == false {
			{{toLower parent}}.Name = in.Name.String
		}
		if in.Desc.IsZero() == false {
			{{toLower parent}}.Desc = in.Desc.String
		}

		err = {{toLower parent}}s.Update(r.Context(), {{toLower parent}})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("name", {{toLower parent}}.Name).
				WithField("id", id).
				Errorln("cannot update the {{toLower parent}}")
		} else {
			render.JSON(w, {{toLower parent}}, 200)
		}
	}
}

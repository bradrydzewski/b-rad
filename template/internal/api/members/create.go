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

package members

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

// HandleCreate returns an http.HandlerFunc that processes
// a request to create user membership for a project.
func HandleCreate(
	users store.UserStore,
	projects store.ProjectStore,
	members store.MemberStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		in := new(types.MembershipInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot unmarshal json request")
			return
		}

		id, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		project, err := projects.Find(r.Context(), id)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("project", project).
				Debugln("project not found")
			return
		}

		key := chi.URLParam(r, "user")
		user, err := users.FindKey(r.Context(), key)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("user", key).
				WithField("project.id", project.ID).
				WithField("project.name", project.Name).
				Debugln("user not found")
			return
		}

		membership := &types.Membership{
			Project: project.ID,
			User:    user.ID,
			Role:    in.Role,
		}
		err = members.Create(r.Context(), membership)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("user.email", user.Email).
				WithField("project.id", project.ID).
				WithField("project.name", project.Name).
				Errorln("cannot create member")
			return
		}

		member := &types.Member{
			Email: user.Email,
			Role:  membership.Role,
		}
		render.JSON(w, member, 200)
	}
}

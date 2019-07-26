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

package projects

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/api/request"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"
	"github.com/{{github}}/types/enum"

	"github.com/dchest/uniuri"
)

// HandleCreate returns an http.HandlerFunc that creates
// a new project.
func HandleCreate(projects store.ProjectStore, members store.MemberStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		viewer, _ := request.UserFrom(ctx)

		in := new(types.ProjectInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot unmarshal json request")
			return
		}

		project := &types.Project{
			Name:    in.Name.String,
			Desc:    in.Desc.String,
			Token:   uniuri.NewLen(uniuri.UUIDLen),
			Created: time.Now().Unix(),
			Updated: time.Now().Unix(),
		}

		err = projects.Create(ctx, project)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("name", in.Name).
				Errorf("cannot create project")
			return
		}

		membership := &types.Membership{
			Project: project.ID,
			User:    viewer.ID,
			Role:    enum.RoleAdmin,
		}
		if err := members.Create(ctx, membership); err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("user.email", viewer.Email).
				WithField("project.id", project.ID).
				WithField("project.name", project.Name).
				Errorln("cannot create default membership")
			return
		}

		render.JSON(w, project, 200)
	}
}

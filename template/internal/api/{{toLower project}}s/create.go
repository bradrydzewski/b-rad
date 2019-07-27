// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower project}}s

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/api/request"
	"github.com/{{toLower repo}}/internal/logger"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/{{toLower repo}}/types/enum"

	"github.com/dchest/uniuri"
)

// HandleCreate returns an http.HandlerFunc that creates
// a new {{toLower project}}.
func HandleCreate({{toLower project}}s store.{{title project}}Store, members store.MemberStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		viewer, _ := request.UserFrom(ctx)

		in := new(types.{{title project}}Input)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot unmarshal json request")
			return
		}

		{{toLower project}} := &types.{{title project}}{
			Name:    in.Name.String,
			Desc:    in.Desc.String,
			Token:   uniuri.NewLen(uniuri.UUIDLen),
			Created: time.Now().Unix(),
			Updated: time.Now().Unix(),
		}

		err = {{toLower project}}s.Create(ctx, {{toLower project}})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("name", in.Name).
				Errorf("cannot create {{toLower project}}")
			return
		}

		membership := &types.Membership{
			{{title project}}: {{toLower project}}.ID,
			User:    viewer.ID,
			Role:    enum.RoleAdmin,
		}
		if err := members.Create(ctx, membership); err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("user.email", viewer.Email).
				WithField("{{toLower project}}.id", {{toLower project}}.ID).
				WithField("{{toLower project}}.name", {{toLower project}}.Name).
				Errorln("cannot create default membership")
			return
		}

		render.JSON(w, {{toLower project}}, 200)
	}
}

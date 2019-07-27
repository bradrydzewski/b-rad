// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package members

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/logger"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"

	"github.com/go-chi/chi"
)

// HandleCreate returns an http.HandlerFunc that processes
// a request to create user membership for a {{toLower project}}.
func HandleCreate(
	users store.UserStore,
	{{toLower project}}s store.{{title project}}Store,
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

		id, err := strconv.ParseInt(chi.URLParam(r, "{{toLower project}}"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse {{toLower project}} id")
			return
		}

		{{toLower project}}, err := {{toLower project}}s.Find(r.Context(), id)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("{{toLower project}}", {{toLower project}}).
				Debugln("{{toLower project}} not found")
			return
		}

		key := chi.URLParam(r, "user")
		user, err := users.FindKey(r.Context(), key)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("user", key).
				WithField("{{toLower project}}.id", {{toLower project}}.ID).
				WithField("{{toLower project}}.name", {{toLower project}}.Name).
				Debugln("user not found")
			return
		}

		membership := &types.Membership{
			{{title project}}: {{toLower project}}.ID,
			User:    user.ID,
			Role:    in.Role,
		}
		err = members.Create(r.Context(), membership)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("user.email", user.Email).
				WithField("{{toLower project}}.id", {{toLower project}}.ID).
				WithField("{{toLower project}}.name", {{toLower project}}.Name).
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

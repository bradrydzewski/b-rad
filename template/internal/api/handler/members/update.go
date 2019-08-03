// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package members

import (
	"encoding/json"
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/rs/zerolog/hlog"

	"github.com/go-chi/chi"
)

// HandleUpdate returns an http.HandlerFunc that processes
// a request to update account membership to a {{toLower project}}.
func HandleUpdate(
	users store.UserStore,
	{{toLower project}}s store.{{title project}}Store,
	members store.MemberStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := hlog.FromRequest(r)

		in := new(types.MembershipInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			log.Debug().Err(err).
				Msg("cannot unmarshal json request")
			return
		}

		slug := chi.URLParam(r, "{{toLower project}}")
		{{toLower project}}, err := {{toLower project}}s.FindSlug(ctx, slug)
		if err != nil {
			render.NotFound(w, err)
			log.Debug().Err(err).
				Str("{{toLower project}}_slug", slug).
				Msg("{{toLower project}} not found")
			return
		}

		email := chi.URLParam(r, "user")
		user, err := users.FindKey(ctx, email)
		if err != nil {
			render.NotFound(w, err)
			log.Debug().Err(err).
				Str("user", email).
				Str("{{toLower project}}_name", {{toLower project}}.Name).
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Msg("user not found")
			return
		}

		membership := &types.Membership{
			{{title project}}: {{toLower project}}.ID,
			User:    user.ID,
			Role:    in.Role,
		}
		err = members.Update(ctx, membership)
		if err != nil {
			render.InternalError(w, err)
			log.Error().Err(err).
				Str("user_email", user.Email).
				Int64("user_id", user.ID).
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Str("{{toLower project}}_name", {{toLower project}}.Name).
				Msg("cannot create member")
			return
		}

		member := &types.Member{
			Email: user.Email,
			Role:  membership.Role,
		}
		render.JSON(w, member, 200)
	}
}

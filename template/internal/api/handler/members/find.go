// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package members

import (
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/rs/zerolog/hlog"

	"github.com/go-chi/chi"
)

// HandleFind returns an http.HandlerFunc that writes a json-encoded
// {{toLower project}} members details to the response body.
func HandleFind(
	users store.UserStore,
	{{toLower project}}s store.{{title project}}Store,
	members store.MemberStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx  = r.Context()
			log  = hlog.FromRequest(r)
			slug = chi.URLParam(r, "{{toLower project}}")
		)

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
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Str("user", email).
				Msg("user not found")
			return
		}
		member, err := members.Find(ctx, {{toLower project}}.ID, user.ID)
		if err != nil {
			render.NotFound(w, err)
			log.Debug().Err(err).
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Str("user_email", user.Email).
				Int64("user_email", user.ID).
				Msg("membership not found")
			return
		}
		render.JSON(w, member, 200)
	}
}

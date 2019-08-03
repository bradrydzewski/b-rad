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
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/{{toLower repo}}/types/check"
	"github.com/{{toLower repo}}/types/enum"

	"github.com/dchest/uniuri"
	"github.com/gosimple/slug"
	"github.com/gotidy/ptr"
	"github.com/rs/zerolog/hlog"
)

// HandleCreate returns an http.HandlerFunc that creates
// a new {{toLower project}}.
func HandleCreate({{toLower project}}s store.{{title project}}Store, members store.MemberStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := hlog.FromRequest(r)
		viewer, _ := request.UserFrom(ctx)

		in := new(types.{{title project}}Input)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			log.Debug().Err(err).
				Msg("cannot unmarshal json request")
			return
		}

		{{toLower project}} := &types.{{title project}}{
			Slug:    ptr.ToString(in.Slug),
			Name:    ptr.ToString(in.Name),
			Desc:    ptr.ToString(in.Desc),
			Token:   uniuri.NewLen(uniuri.UUIDLen),
			Created: time.Now().Unix(),
			Updated: time.Now().Unix(),
		}

		// if the slug is empty we can derrive
		// the slug from the {{toLower project}} name.
		if {{toLower project}}.Slug == "" {
			{{toLower project}}.Slug = slug.Make({{toLower project}}.Name)
		}

		// if the name is empty we can derrive
		// the name from the {{toLower project}} slug.
		if {{toLower project}}.Name == "" {
			{{toLower project}}.Name = {{toLower project}}.Slug
		}

		if ok, err := check.{{title project}}({{toLower project}}); !ok {
			render.BadRequest(w, err)
			log.Debug().Err(err).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Msg("cannot create {{toLower project}}")
			return
		}

		err = {{toLower project}}s.Create(ctx, {{toLower project}})
		if err != nil {
			render.InternalError(w, err)
			log.Error().Err(err).
				Str("{{toLower project}}_name", {{toLower project}}.Name).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Msg("cannot create {{toLower project}}")
			return
		}

		membership := &types.Membership{
			{{title project}}: {{toLower project}}.ID,
			User:    viewer.ID,
			Role:    enum.RoleAdmin,
		}
		if err := members.Create(ctx, membership); err != nil {
			render.InternalError(w, err)
			log.Error().Err(err).
				Int64("user_id", viewer.ID).
				Str("user_email", viewer.Email).
				Str("{{toLower project}}_name", {{toLower project}}.Name).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Msg("cannot create default membership")
			return
		}

		render.JSON(w, {{toLower project}}, 200)
	}
}

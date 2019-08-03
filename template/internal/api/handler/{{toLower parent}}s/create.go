// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower parent}}s

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/{{toLower repo}}/types/check"
	"github.com/go-chi/chi"
	"github.com/gosimple/slug"
	"github.com/gotidy/ptr"
	"github.com/rs/zerolog/hlog"
)

// HandleCreate returns an http.HandlerFunc that creates
// the object and persists to the datastore.
func HandleCreate({{toLower project}}s store.{{title project}}Store, {{toLower parent}}s store.{{title parent}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx   = r.Context()
			log   = hlog.FromRequest(r)
			param = chi.URLParam(r, "{{toLower project}}")
		)

		{{toLower project}}, err := {{toLower project}}s.FindSlug(ctx, param)
		if err != nil {
			render.NotFound(w, err)
			log.Debug().Err(err).
				Str("{{toLower project}}_slug", param).
				Msg("{{toLower project}} not found")
			return
		}

		sublog := log.With().
			Int64("{{toLower project}}_id", {{toLower project}}.ID).
			Str("{{toLower project}}_slug", {{toLower project}}.Slug).
			Logger()

		in := new(types.{{title parent}}Input)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			sublog.Debug().Err(err).
				Msg("cannot unmarshal json request")
			return
		}

		{{toLower parent}} := &types.{{title parent}}{
			{{title project}}: {{toLower project}}.ID,
			Slug:    ptr.ToString(in.Slug),
			Name:    ptr.ToString(in.Name),
			Desc:    ptr.ToString(in.Desc),
			Created: time.Now().Unix(),
			Updated: time.Now().Unix(),
		}

		// if the slug is empty we can derrive
		// the slug from the name.
		if {{toLower parent}}.Slug == "" {
			{{toLower parent}}.Slug = slug.Make({{toLower parent}}.Name)
		}

		// if the name is empty we can derrive
		// the name from the slug.
		if {{toLower parent}}.Name == "" {
			{{toLower parent}}.Name = {{toLower parent}}.Slug
		}

		if ok, err := check.{{title parent}}({{toLower parent}}); !ok {
			render.BadRequest(w, err)
			sublog.Debug().Err(err).
				Int64("{{toLower parent}}_id", {{toLower parent}}.ID).
				Str("{{toLower parent}}_slug", {{toLower parent}}.Slug).
				Msg("cannot validate {{toLower parent}}")
			return
		}

		err = {{toLower parent}}s.Create(ctx, {{toLower parent}})
		if err != nil {
			render.InternalError(w, err)
			sublog.Error().Err(err).
				Int64("{{toLower parent}}_id", {{toLower parent}}.ID).
				Str("{{toLower parent}}_slug", {{toLower parent}}.Slug).
				Msg("cannot create {{toLower parent}}")
		} else {
			render.JSON(w, {{toLower parent}}, 200)
		}
	}
}

// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower child}}s

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
func HandleCreate({{toLower project}}s store.{{title project}}Store, {{toLower parent}}s store.{{title parent}}Store, {{toLower child}}s store.{{title child}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx          = r.Context()
			log          = hlog.FromRequest(r)
			{{toLower project}}Param = chi.URLParam(r, "{{toLower project}}")
			{{toLower parent}}Param   = chi.URLParam(r, "{{toLower parent}}")
		)

		{{toLower project}}, err := {{toLower project}}s.FindSlug(ctx, {{toLower project}}Param)
		if err != nil {
			render.NotFound(w, err)
			log.Debug().Err(err).
				Str("{{toLower project}}_slug", {{toLower project}}Param).
				Msg("{{toLower project}} not found")
			return
		}

		sublog := log.With().
			Int64("{{toLower project}}_id", {{toLower project}}.ID).
			Str("{{toLower project}}_slug", {{toLower project}}.Slug).
			Logger()

		{{toLower parent}}, err := {{toLower parent}}s.FindSlug(ctx, {{toLower project}}.ID, {{toLower parent}}Param)
		if err != nil {
			render.NotFound(w, err)
			sublog.Debug().Err(err).
				Str("{{toLower parent}}_slug", {{toLower parent}}Param).
				Msg("{{toLower parent}} not found")
			return
		}

		sublog = sublog.With().
			Int64("{{toLower parent}}_id", {{toLower parent}}.ID).
			Str("{{toLower parent}}_slug", {{toLower parent}}.Slug).
			Logger()

		in := new(types.{{title child}}Input)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			sublog.Debug().Err(err).
				Msg("cannot unmarshal json request")
			return
		}

		{{toLower child}} := &types.{{title child}}{
			{{title project}}: {{toLower project}}.ID,
			{{title parent}}:   {{toLower parent}}.ID,
			Slug:    ptr.ToString(in.Slug),
			Name:    ptr.ToString(in.Name),
			Desc:    ptr.ToString(in.Desc),
			Created: time.Now().Unix(),
			Updated: time.Now().Unix(),
		}

		// if the slug is empty we can derrive
		// the slug from the name.
		if {{toLower child}}.Slug == "" {
			{{toLower child}}.Slug = slug.Make({{toLower child}}.Name)
		}

		// if the name is empty we can derrive
		// the name from the slug.
		if {{toLower child}}.Name == "" {
			{{toLower child}}.Name = {{toLower child}}.Slug
		}

		if ok, err := check.{{title child}}({{toLower child}}); !ok {
			render.BadRequest(w, err)
			sublog.Debug().Err(err).
				Str("{{toLower child}}_slug", {{toLower child}}.Slug).
				Msg("cannot create {{toLower child}}")
			return
		}

		err = {{toLower child}}s.Create(ctx, {{toLower child}})
		if err != nil {
			render.InternalError(w, err)
			sublog.Error().Err(err).
				Str("{{toLower child}}_slug", {{toLower child}}.Slug).
				Msg("cannot create {{toLower child}}")
		} else {
			render.JSON(w, {{toLower child}}, 200)
		}
	}
}

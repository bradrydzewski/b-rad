// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower parent}}s

import (
	"encoding/json"
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/{{toLower repo}}/types/check"

	"github.com/go-chi/chi"
	"github.com/gotidy/ptr"
	"github.com/rs/zerolog/hlog"
)

// HandleUpdate returns an http.HandlerFunc that processes http
// requests to update the object details.
func HandleUpdate({{toLower project}}s store.{{title project}}Store, {{toLower parent}}s store.{{title parent}}Store) http.HandlerFunc {
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
			Str("{{toLower parent}}_slug", {{toLower parent}}.Slug).
			Int64("{{toLower parent}}_id", {{toLower parent}}.ID).
			Logger()

		in := new(types.{{title parent}}Input)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			sublog.Debug().Err(err).
				Msg("cannot unmarshal json request")
			return
		}

		if in.Name != nil {
			{{toLower parent}}.Name = ptr.ToString(in.Name)
		}
		if in.Desc != nil {
			{{toLower parent}}.Desc = ptr.ToString(in.Desc)
		}

		if ok, err := check.{{title parent}}({{toLower parent}}); !ok {
			render.BadRequest(w, err)
			sublog.Debug().Err(err).
				Msg("cannot validate {{toLower parent}}")
			return
		}

		err = {{toLower parent}}s.Update(ctx, {{toLower parent}})
		if err != nil {
			render.InternalError(w, err)
			sublog.Error().Err(err).
				Msg("cannot update {{toLower parent}}")
		} else {
			render.JSON(w, {{toLower parent}}, 200)
		}
	}
}

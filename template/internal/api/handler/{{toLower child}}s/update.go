// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower child}}s

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
func HandleUpdate({{toLower project}}s store.{{title project}}Store, {{toLower parent}}s store.{{title parent}}Store, {{toLower child}}s store.{{title child}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx          = r.Context()
			log          = hlog.FromRequest(r)
			{{toLower project}}Param = chi.URLParam(r, "{{toLower project}}")
			{{toLower parent}}Param   = chi.URLParam(r, "{{toLower parent}}")
			{{toLower child}}Param = chi.URLParam(r, "{{toLower child}}")
		)

		// find the {{toLower project}} using the {{toLower project}} slug from the
		// url path parameter.
		{{toLower project}}, err := {{toLower project}}s.FindSlug(ctx, {{toLower project}}Param)
		if err != nil {
			render.NotFound(w, err)
			log.Debug().Err(err).
				Str("{{toLower project}}_slug", {{toLower project}}Param).
				Msg("{{toLower project}} not found")
			return
		}

		// create sublog with {{toLower project}} metadata
		sublog := log.With().
			Int64("{{toLower project}}_id", {{toLower project}}.ID).
			Str("{{toLower project}}_slug", {{toLower project}}.Slug).
			Logger()

		// find the {{toLower parent}} using the slug from the url
		// path parameter.
		{{toLower parent}}, err := {{toLower parent}}s.FindSlug(ctx, {{toLower project}}.ID, {{toLower parent}}Param)
		if err != nil {
			render.NotFound(w, err)
			sublog.Debug().Err(err).
				Str("{{toLower parent}}_slug", {{toLower parent}}Param).
				Msg("{{toLower parent}} not found")
			return
		}

		// create sublog with {{toLower parent}} metadata
		sublog = sublog.With().
			Int64("{{toLower parent}}_id", {{toLower parent}}.ID).
			Str("{{toLower parent}}_slug", {{toLower parent}}.Slug).
			Logger()

		// find the {{toLower child}} using the slug from the
		// url path parameter.
		{{toLower child}}, err := {{toLower child}}s.FindSlug(ctx, {{toLower parent}}.ID, {{toLower child}}Param)
		if err != nil {
			render.NotFound(w, err)
			sublog.Debug().Err(err).
				Str("{{toLower child}}_slug", {{toLower child}}Param).
				Msg("{{toLower child}} not found")
			return
		}

		// create sublog with {{toLower child}} metadata
		sublog = sublog.With().
			Int64("{{toLower child}}_id", {{toLower child}}.ID).
			Str("{{toLower child}}_slug", {{toLower child}}.Slug).
			Logger()

		// unmarshal the json input provided in the
		// http.Request body.
		in := new(types.{{title child}}Input)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			sublog.Debug().Err(err).
				Msg("cannot unmarshal json request")
			return
		}

		if in.Name != nil {
			{{toLower child}}.Name = ptr.ToString(in.Name)
		}
		if in.Desc != nil {
			{{toLower child}}.Desc = ptr.ToString(in.Desc)
		}

		// check to ensure the {{toLower child}} is still valid
		// after user-provided updates are applied.
		if ok, err := check.{{title child}}({{toLower child}}); !ok {
			render.BadRequest(w, err)
			sublog.Debug().Err(err).
				Msg("cannot update {{toLower child}}")
			return
		}

		// update the {{toLower child}} in the datastore.
		err = {{toLower child}}s.Update(ctx, {{toLower child}})
		if err != nil {
			render.InternalError(w, err)
			sublog.Error().Err(err).
				Msg("cannot update {{toLower child}}")
		} else {
			render.JSON(w, {{toLower child}}, 200)
		}
	}
}

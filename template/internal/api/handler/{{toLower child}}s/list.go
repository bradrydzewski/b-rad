// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower child}}s

import (
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/api/request"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog/hlog"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of objects to the response body.
func HandleList({{toLower project}}s store.{{title project}}Store, {{toLower parent}}s store.{{title parent}}Store, {{toLower child}}s store.{{title child}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx          = r.Context()
			log          = hlog.FromRequest(r)
			page         = request.ParsePage(r)
			size         = request.ParseSize(r)
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

		items, err := {{toLower child}}s.List(ctx, {{toLower parent}}.ID, types.Params{
			Size: size,
			Page: page,
		})
		if err != nil {
			render.InternalError(w, err)
			sublog.Error().Err(err).
				Str("{{toLower parent}}_slug", {{toLower parent}}.Slug).
				Int64("{{toLower parent}}_id", {{toLower parent}}.ID).
				Msg("cannot retrieve list")
		} else {
			render.Pagination(r, w, page, size, 0)
			render.JSON(w, items, 200)
		}
	}
}

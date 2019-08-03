// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower parent}}s

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
func HandleList({{toLower project}}s store.{{title project}}Store, {{toLower parent}}s store.{{title parent}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx  = r.Context()
			log  = hlog.FromRequest(r)
			slug = chi.URLParam(r, "{{toLower project}}")
			page = request.ParsePage(r)
			size = request.ParseSize(r)
		)

		{{toLower project}}, err := {{toLower project}}s.FindSlug(ctx, slug)
		if err != nil {
			render.NotFound(w, err)
			log.Debug().Err(err).
				Str("{{toLower project}}_slug", slug).
				Msg("{{toLower project}} not found")
			return
		}

		{{toLower parent}}s, err := {{toLower parent}}s.List(ctx, {{toLower project}}.ID, types.Params{
			Size: size,
			Page: page,
		})
		if err != nil {
			render.InternalError(w, err)
			log.Error().Err(err).
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Msg("cannot retrieve list")
		} else {
			render.Pagination(r, w, page, size, 0)
			render.JSON(w, {{toLower parent}}s, 200)
		}
	}
}

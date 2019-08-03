// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower parent}}s

import (
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/store"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/hlog"
)

// HandleDelete returns an http.HandlerFunc that deletes
// the object from the datastore.
func HandleDelete({{toLower project}}s store.{{title project}}Store, {{toLower parent}}s store.{{title parent}}Store) http.HandlerFunc {
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

		{{toLower parent}}, err := {{toLower parent}}s.FindSlug(ctx, {{toLower project}}.ID, {{toLower parent}}Param)
		if err != nil {
			render.NotFound(w, err)
			log.Debug().Err(err).
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Str("{{toLower parent}}_slug", {{toLower parent}}Param).
				Msg("{{toLower parent}} not found")
			return
		}

		err = {{toLower parent}}s.Delete(ctx, {{toLower parent}})
		if err != nil {
			render.InternalError(w, err)
			log.Error().Err(err).
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Int64("{{toLower parent}}_id", {{toLower parent}}.ID).
				Str("{{toLower parent}}_slug", {{toLower parent}}.Slug).
				Msg("cannot delete {{toLower parent}}")
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

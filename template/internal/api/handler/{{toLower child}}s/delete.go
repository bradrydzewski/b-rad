// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower child}}s

import (
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/store"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/hlog"
)

// HandleDelete returns an http.HandlerFunc that deletes
// the object from the datastore.
func HandleDelete({{toLower project}}s store.{{title project}}Store, {{toLower parent}}s store.{{title parent}}Store, {{toLower child}}s store.{{title child}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx          = r.Context()
			log          = hlog.FromRequest(r)
			{{toLower project}}Param = chi.URLParam(r, "{{toLower project}}")
			{{toLower parent}}Param   = chi.URLParam(r, "{{toLower parent}}")
			{{toLower child}}Param = chi.URLParam(r, "{{toLower child}}")
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

		{{toLower child}}, err := {{toLower child}}s.FindSlug(ctx, {{toLower parent}}.ID, {{toLower child}}Param)
		if err != nil {
			render.NotFound(w, err)
			sublog.Debug().Err(err).
				Str("{{toLower child}}_slug", {{toLower child}}Param).
				Msg("{{toLower parent}} not found")
			return
		}

		err = {{toLower child}}s.Delete(ctx, {{toLower child}})
		if err != nil {
			render.InternalError(w, err)
			sublog.Error().Err(err).
				Int64("{{toLower child}}_id", {{toLower child}}.ID).
				Str("{{toLower child}}_slug", {{toLower child}}.Slug).
				Msg("cannot delete {{toLower child}}")
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

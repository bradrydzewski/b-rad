// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower project}}s

import (
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/store"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/hlog"
)

// HandleDelete returns an http.HandlerFunc that deletes
// the object from the datastore.
func HandleDelete({{toLower project}}s store.{{title project}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "{{toLower project}}")

		{{toLower project}}, err := {{toLower project}}s.FindSlug(ctx, id)
		if err != nil {
			render.NotFound(w, err)
			hlog.FromRequest(r).
				Debug().Err(err).
				Str("{{toLower project}}_slug", id).
				Msg("{{toLower project}} not found")
			return
		}

		err = {{toLower project}}s.Delete(ctx, {{toLower project}})
		if err != nil {
			render.InternalError(w, err)
			hlog.FromRequest(r).
				Error().Err(err).
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Msg("cannot delete {{toLower project}}")
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower project}}s

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
// requests to update the {{toLower project}} details.
func HandleUpdate({{toLower project}}s store.{{title project}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := hlog.FromRequest(r)
		id := chi.URLParam(r, "{{toLower project}}")

		{{toLower project}}, err := {{toLower project}}s.FindSlug(ctx, id)
		if err != nil {
			render.NotFound(w, err)
			log.Debug().Err(err).
				Str("{{toLower project}}_slug", id).
				Msg("{{toLower project}} not found")
			return
		}

		in := new(types.{{title project}}Input)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			log.Debug().Err(err).
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Msg("cannot unmarshal json request")
			return
		}

		if in.Name != nil {
			{{toLower project}}.Name = ptr.ToString(in.Name)
		}
		if in.Desc != nil {
			{{toLower project}}.Desc = ptr.ToString(in.Desc)
		}

		if ok, err := check.{{title project}}({{toLower project}}); !ok {
			render.BadRequest(w, err)
			log.Debug().Err(err).
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Msg("cannot update {{toLower project}}")
			return
		}

		err = {{toLower project}}s.Update(ctx, {{toLower project}})
		if err != nil {
			render.InternalError(w, err)
			log.Error().Err(err).
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Msg("cannot update the {{toLower project}}")
		} else {
			render.JSON(w, {{toLower project}}, 200)
		}
	}
}

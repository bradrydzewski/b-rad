// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package members

import (
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/api/request"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/hlog"
)

// HandleList returns an http.HandlerFunc that write a json-encoded
// list of {{toLower project}} members to the response body.
func HandleList(
	{{toLower project}}s store.{{title project}}Store,
	members store.MemberStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			slug = chi.URLParam(r, "{{toLower project}}")
			page = request.ParsePage(r)
			size = request.ParseSize(r)
			ctx  = r.Context()
			log  = hlog.FromRequest(r)
		)

		{{toLower project}}, err := {{toLower project}}s.FindSlug(ctx, slug)
		if err != nil {
			render.NotFound(w, err)
			log.Debug().Err(err).
				Str("{{toLower project}}_slug", slug).
				Msg("{{toLower project}} not found")
			return
		}
		members, err := members.List(ctx, {{toLower project}}.ID, types.Params{
			Size: size,
			Page: page,
		})
		if err != nil {
			render.InternalError(w, err)
			log.Error().Err(err).
				Int64("{{toLower project}}_id", {{toLower project}}.ID).
				Str("{{toLower project}}_slug", {{toLower project}}.Slug).
				Msg("cannot list members")
		} else {
			render.Pagination(r, w, page, size, 0)
			render.JSON(w, members, 200)
		}
	}
}

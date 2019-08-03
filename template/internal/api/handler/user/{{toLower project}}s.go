// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package user

import (
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/api/request"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"

	"github.com/rs/zerolog/log"
)

// Handle{{title project}}s returns an http.HandlerFunc that writes a json-encoded
// list of {{toLower project}}s to the response body.
func Handle{{title project}}s({{toLower project}}s store.{{title project}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			page = request.ParsePage(r)
			size = request.ParseSize(r)
			ctx  = r.Context()
		)

		viewer, _ := request.UserFrom(ctx)
		list, err := {{toLower project}}s.List(ctx, viewer.ID, types.Params{
			Page: page,
			Size: size,
		})
		if err != nil {
			render.InternalError(w, err)
			log.Ctx(ctx).Error().
				Err(err).Msg("cannot list {{toLower project}}s")
		} else {
			render.Pagination(r, w, page, size, 0)
			render.JSON(w, list, 200)
		}
	}
}

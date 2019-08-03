// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package users

import (
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/api/request"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types/enum"
	"github.com/rs/zerolog/hlog"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of all registered system users to the response body.
func HandleList(users store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = r.Context()
			log = hlog.FromRequest(r)
		)

		params := request.ParseUserFilter(r)
		if params.Order == enum.OrderDefault {
			params.Order = enum.OrderAsc
		}

		count, err := users.Count(ctx)
		if err != nil {
			log.Error().Err(err).
				Msg("cannot retrieve user count")
		}

		list, err := users.List(ctx, params)
		if err != nil {
			render.InternalError(w, err)
			log.Error().Err(err).
				Msg("cannot retrieve user list")
			return
		}

		render.Pagination(r, w, params.Page, params.Size, int(count))
		render.JSON(w, list, 200)

	}
}

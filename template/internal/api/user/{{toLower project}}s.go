// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package user

import (
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/api/request"
	"github.com/{{toLower repo}}/internal/logger"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
)

// Handle{{title project}}s returns an http.HandlerFunc that writes a json-encoded
// list of {{toLower project}}s to the response body.
func Handle{{title project}}s({{toLower project}}s store.{{title project}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())
		list, err := {{toLower project}}s.List(r.Context(), viewer.ID, types.Params{})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("email", viewer.Email).
				Errorf("cannot list {{toLower project}}s")
		} else {
			render.JSON(w, list, 200)
		}
	}
}

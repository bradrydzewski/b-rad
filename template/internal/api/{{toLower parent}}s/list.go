// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower parent}}s

import (
	"net/http"
	"strconv"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/logger"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/go-chi/chi"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of objects to the response body.
func HandleList({{toLower parent}}s store.{{title parent}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := strconv.ParseInt(chi.URLParam(r, "{{toLower project}}"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse {{toLower project}} id")
			return
		}

		{{toLower parent}}s, err := {{toLower parent}}s.List(ctx, id, types.Params{})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				Errorf("cannot retrieve list")
		} else {
			render.JSON(w, {{toLower parent}}s, 200)
		}
	}
}

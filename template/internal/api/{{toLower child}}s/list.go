// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower child}}s

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
func HandleList({{toLower parent}}s store.{{title parent}}Store, {{toLower child}}s store.{{title child}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		{{toLower project}}, err := strconv.ParseInt(chi.URLParam(r, "{{toLower project}}"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse {{toLower project}} id")
			return
		}

		id, err := strconv.ParseInt(chi.URLParam(r, "{{toLower parent}}"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse {{toLower parent}} id")
			return
		}

		{{toLower parent}}, err := {{toLower parent}}s.Find(ctx, id)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				Errorf("cannot retrieve {{toLower parent}}")
			return
		}

		if {{toLower parent}}.{{title project}} != {{toLower project}} {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithError(err).
				Errorf("mismatch {{toLower project}} id")
			return
		}

		items, err := {{toLower child}}s.List(ctx, {{toLower parent}}.ID, types.Params{})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				Errorf("cannot retrieve list")
		} else {
			render.JSON(w, items, 200)
		}
	}
}

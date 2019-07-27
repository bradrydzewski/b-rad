// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower project}}s

import (
	"net/http"
	"strconv"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/logger"
	"github.com/{{toLower repo}}/internal/store"

	"github.com/go-chi/chi"
)

// HandleDelete returns an http.HandlerFunc that deletes
// the object from the datastore.
func HandleDelete({{toLower project}}s store.{{title project}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "{{toLower project}}"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse {{toLower project}} id")
			return
		}

		{{toLower project}}, err := {{toLower project}}s.Find(r.Context(), id)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", id).
				Debugln("{{toLower project}} not found")
			return
		}

		err = {{toLower project}}s.Delete(r.Context(), {{toLower project}})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", id).
				Debugln("cannot delete {{toLower project}}")
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

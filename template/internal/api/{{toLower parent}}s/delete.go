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

	"github.com/go-chi/chi"
)

// HandleDelete returns an http.HandlerFunc that deletes
// the object from the datastore.
func HandleDelete({{toLower parent}}s store.{{title parent}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		{{toLower parent}}, err := {{toLower parent}}s.Find(r.Context(), id)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", id).
				Debugln("{{toLower parent}} not found")
			return
		}

		if {{toLower parent}}.{{title project}} != {{toLower project}} {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithField("id", id).
				WithField("{{toLower project}}", {{toLower project}}).
				Debugln("{{toLower project}} id mismatch")
			return
		}

		err = {{toLower parent}}s.Delete(r.Context(), {{toLower parent}})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", id).
				Debugln("cannot delete {{toLower parent}}")
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

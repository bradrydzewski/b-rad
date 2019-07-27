// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower parent}}s

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/logger"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"

	"github.com/go-chi/chi"
)

// HandleUpdate returns an http.HandlerFunc that processes http
// requests to update the object details.
func HandleUpdate({{toLower parent}}s store.{{title parent}}Store) http.HandlerFunc {
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

		in := new(types.{{title parent}}Input)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("{{toLower project}}", {{toLower project}}).
				WithField("id", id).
				Debugln("cannot unmarshal json request")
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

		if in.Name.IsZero() == false {
			{{toLower parent}}.Name = in.Name.String
		}
		if in.Desc.IsZero() == false {
			{{toLower parent}}.Desc = in.Desc.String
		}

		err = {{toLower parent}}s.Update(r.Context(), {{toLower parent}})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("name", {{toLower parent}}.Name).
				WithField("id", id).
				Errorln("cannot update the {{toLower parent}}")
		} else {
			render.JSON(w, {{toLower parent}}, 200)
		}
	}
}

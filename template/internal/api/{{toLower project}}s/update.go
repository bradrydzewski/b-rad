// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower project}}s

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
// requests to update the {{toLower project}} details.
func HandleUpdate({{toLower project}}s store.{{title project}}Store) http.HandlerFunc {
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

		in := new(types.{{title project}}Input)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("name", {{toLower project}}.Name).
				WithField("id", id).
				Debugln("cannot unmarshal json request")
			return
		}

		if in.Name.IsZero() == false {
			{{toLower project}}.Name = in.Name.String
		}
		if in.Desc.IsZero() == false {
			{{toLower project}}.Desc = in.Desc.String
		}

		err = {{toLower project}}s.Update(r.Context(), {{toLower project}})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("name", {{toLower project}}.Name).
				WithField("id", id).
				Errorln("cannot update the {{toLower project}}")
		} else {
			render.JSON(w, {{toLower project}}, 200)
		}
	}
}

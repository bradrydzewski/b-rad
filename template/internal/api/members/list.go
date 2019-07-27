// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package members

import (
	"net/http"
	"strconv"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/logger"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"

	"github.com/go-chi/chi"
)

// HandleList returns an http.HandlerFunc that write a json-encoded
// list of {{toLower project}} members to the response body.
func HandleList(
	{{toLower project}}s store.{{title project}}Store,
	members store.MemberStore,
) http.HandlerFunc {
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
				WithField("{{toLower project}}", id).
				Debugln("{{toLower project}} not found")
			return
		}
		members, err := members.List(r.Context(), {{toLower project}}.ID, types.Params{})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("{{toLower project}}", {{toLower project}}.ID).
				WithField("name", {{toLower project}}.Name).
				Errorln("cannot list members")
		} else {
			render.JSON(w, members, 200)
		}
	}
}

// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower project}}s

import (
	"net/http"
	"strconv"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/api/request"
	"github.com/{{toLower repo}}/internal/logger"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/{{toLower repo}}/types/enum"

	"github.com/go-chi/chi"
)

type {{toLower project}}Token struct {
	*types.{{title project}}
	Token string `json:"token"`
}

// HandleFind returns an http.HandlerFunc that writes the
// json-encoded {{toLower project}} details to the response body.
func HandleFind({{toLower project}}s store.{{title project}}Store, members store.MemberStore) http.HandlerFunc {
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

		{{toLower project}}, err := {{toLower project}}s.Find(ctx, id)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", id).
				Debugln("{{toLower project}} not found")
			return
		}

		// if the caller requests the {{toLower project}} details without
		// the token, return the token object as-is.
		if r.FormValue("token") != "true" {
			render.JSON(w, {{toLower project}}, 200)
			return
		}

		// if the caller requests the {{toLower project}} details with
		// the token, verify the user has admin access to
		// the {{toLower project}}.

		viewer, _ := request.UserFrom(ctx)
		member, err := members.Find(ctx, {{toLower project}}.ID, viewer.ID)
		if err != nil || member.Role != enum.RoleAdmin {
			// if the user does not have admin access to the
			// {{toLower project}}, return the {{toLower project}} details without
			// the token.
			render.JSON(w, {{toLower project}}, 200)
			return
		}

		// else if the user has admin access to the {{toLower project}}
		// it is safe to return the token.
		render.JSON(w, &{{toLower project}}Token{{`{`}}{{toLower project}}, {{toLower project}}.Token}, 200)
	}
}

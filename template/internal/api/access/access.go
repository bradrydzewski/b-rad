// Copyright 2019 Brad Rydzewski. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package access

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/api/request"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types/enum"

	"github.com/go-chi/chi"
)

// ProjectAccess returns an http.HandlerFunc middleware that authorizes
// the user read access to the project.
func ProjectAccess(members store.MemberStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			user, ok := request.UserFrom(ctx)
			if !ok {
				render.ErrorCode(w, errors.New("Requires authentication"), 401)
				return
			}

			// if the user is an administrator they are automatically
			// granted access to the endpoint.
			if user.Admin {
				logger.FromRequest(r).
					Debugln("admin user granted read access")
				next.ServeHTTP(w, r)
				return
			}

			id, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
			if err != nil {
				render.BadRequest(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot parse project id")
				return
			}

			member, err := members.Find(ctx, id, user.ID)
			if err != nil {
				render.NotFound(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot find project membership")
				return
			}

			logger.FromRequest(r).
				WithField("role", member.Role).
				Debugln("user granted read access")

			next.ServeHTTP(w, r)
		})
	}
}

// ProjectAdmin returns an http.HandlerFunc middleware that authorizes
// the user admin access to the project.
func ProjectAdmin(members store.MemberStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			user, ok := request.UserFrom(ctx)
			if !ok {
				render.ErrorCode(w, errors.New("Requires authentication"), 401)
				return
			}

			// if the user is an administrator they are automatically
			// granted access to the endpoint.
			if user.Admin {
				next.ServeHTTP(w, r)
				return
			}

			id, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
			if err != nil {
				render.BadRequest(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot parse project id")
				return
			}

			member, err := members.Find(ctx, id, user.ID)
			if err != nil {
				render.NotFound(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot find project membership")
				return
			}

			if member.Role != enum.RoleAdmin {
				render.ErrorCode(w, errors.New("Forbidden"), 403)
				logger.FromRequest(r).
					WithError(err).
					Debugln("insufficient privileges")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// SystemAdmin returns an http.HandlerFunc middleware that authorizes
// the user access to system administration capabilities.
func SystemAdmin() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			user, ok := request.UserFrom(ctx)
			if !ok {
				render.ErrorCode(w, errors.New("Requires authentication"), 401)
				return
			}
			if !user.Admin {
				render.ErrorCode(w, errors.New("Forbidden"), 403)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package access

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/api/request"
	"github.com/{{toLower repo}}/internal/logger"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types/enum"

	"github.com/go-chi/chi"
)

// {{title project}}Access returns an http.HandlerFunc middleware that authorizes
// the user read access to the {{toLower project}}.
func {{title project}}Access(members store.MemberStore) func(http.Handler) http.Handler {
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

			id, err := strconv.ParseInt(chi.URLParam(r, "{{toLower project}}"), 10, 64)
			if err != nil {
				render.BadRequest(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot parse {{toLower project}} id")
				return
			}

			member, err := members.Find(ctx, id, user.ID)
			if err != nil {
				render.NotFound(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot find {{toLower project}} membership")
				return
			}

			logger.FromRequest(r).
				WithField("role", member.Role).
				Debugln("user granted read access")

			next.ServeHTTP(w, r)
		})
	}
}

// {{title project}}Admin returns an http.HandlerFunc middleware that authorizes
// the user admin access to the {{toLower project}}.
func {{title project}}Admin(members store.MemberStore) func(http.Handler) http.Handler {
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

			id, err := strconv.ParseInt(chi.URLParam(r, "{{toLower project}}"), 10, 64)
			if err != nil {
				render.BadRequest(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot parse {{toLower project}} id")
				return
			}

			member, err := members.Find(ctx, id, user.ID)
			if err != nil {
				render.NotFound(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot find {{toLower project}} membership")
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

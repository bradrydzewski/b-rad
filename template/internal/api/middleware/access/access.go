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
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types/enum"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/hlog"
)

// {{title project}}Access returns an http.HandlerFunc middleware that authorizes
// the user read access to the {{toLower project}}.
func {{title project}}Access({{toLower project}}s store.{{title project}}Store, members store.MemberStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := hlog.FromRequest(r)

			user, ok := request.UserFrom(ctx)
			if !ok {
				render.ErrorCode(w, errors.New("Requires authentication"), 401)
				return
			}

			// if the user is an administrator they are automatically
			// granted access to the endpoint.
			if user.Admin {
				log.Debug().Msg("admin user granted read access")
				next.ServeHTTP(w, r)
				return
			}

			slug := chi.URLParam(r, "{{toLower project}}")
			{{toLower project}}, err := {{toLower project}}s.FindSlug(ctx, slug)
			if err != nil {
				render.NotFound(w, err)
				log.Debug().Err(err).
					Str("{{toLower project}}_slug", slug).
					Msg("cannot find {{toLower project}}")
				return
			}

			member, err := members.Find(ctx, {{toLower project}}.ID, user.ID)
			if err != nil {
				render.NotFound(w, err)
				log.Debug().Err(err).
					Int64("{{toLower project}}_id", {{toLower project}}.ID).
					Str("{{toLower project}}_slug", {{toLower project}}.Slug).
					Msg("cannot find {{toLower project}} membership")
				return
			}

			log.Debug().
				Stringer("role", member.Role).
				Msg("user granted read access")

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
			log := hlog.FromRequest(r)

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
				log.Debug().Err(err).
					Msg("cannot parse {{toLower project}} id")
				return
			}

			member, err := members.Find(ctx, id, user.ID)
			if err != nil {
				render.NotFound(w, err)
				log.Debug().Err(err).
					Msg("cannot find {{toLower project}} membership")
				return
			}

			if member.Role != enum.RoleAdmin {
				render.ErrorCode(w, errors.New("Forbidden"), 403)
				log.Debug().Err(err).
					Msg("insufficient privileges")
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

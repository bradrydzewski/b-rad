// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package account

import (
	"net/http"
	"time"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/{{toLower repo}}/types/check"

	"github.com/dchest/uniuri"
	"github.com/rs/zerolog/hlog"
	"golang.org/x/crypto/bcrypt"
)

// HandleRegister returns an http.HandlerFunc that processes an http.Request
// to register the named user account with the system.
func HandleRegister(users store.UserStore, system store.SystemStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := hlog.FromRequest(r)

		username := r.FormValue("username")
		password := r.FormValue("password")

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			render.InternalError(w, err)
			log.Debug().Err(err).
				Str("email", username).
				Msg("cannot hash password")
			return
		}

		user := &types.User{
			Email:    username,
			Password: string(hash),
			Salt:     uniuri.NewLen(uniuri.UUIDLen),
			Created:  time.Now().Unix(),
			Updated:  time.Now().Unix(),
		}

		if ok, err := check.User(user); !ok {
			render.BadRequest(w, err)
			log.Debug().Err(err).
				Str("email", username).
				Msg("invalid user input")
			return
		}

		if err := users.Create(ctx, user); err != nil {
			render.InternalError(w, err)
			log.Error().Err(err).
				Str("email", username).
				Msg("cannot create user")
			return
		}

		// if the registered user is the first user of the system,
		// assume they are the system administrator and grant the
		// user system admin access.
		if user.ID == 1 {
			user.Admin = true
			if err := users.Update(ctx, user); err != nil {
				log.Error().Err(err).
					Str("email", username).
					Msg("cannot enable admin user")
			}
		}

		expires := time.Now().Add(system.Config(ctx).Token.Expire)
		token, err := generate(user.ID, expires.Unix(), user.Salt)
		if err != nil {
			render.InternalErrorf(w, "Failed to create session")
			log.Error().Err(err).
				Str("email", username).
				Msg("failed to generate token")
			return
		}

		render.JSON(w, &types.UserToken{
			User: user,
			Token: &types.Token{
				Value:   token,
				Expires: expires.UTC(),
			},
		}, 200)
	}
}

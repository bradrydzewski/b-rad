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

	"github.com/rs/zerolog/hlog"
	"golang.org/x/crypto/bcrypt"
)

// HandleLogin returns an http.HandlerFunc that authenticates
// the user and returns an authentication token on success.
func HandleLogin(users store.UserStore, system store.SystemStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := hlog.FromRequest(r)

		username := r.FormValue("username")
		password := r.FormValue("password")
		user, err := users.FindEmail(ctx, username)
		if err != nil {
			render.NotFoundf(w, "Invalid email or password")
			log.Debug().Err(err).
				Str("user", username).
				Msg("cannot find user")
			return
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(user.Password),
			[]byte(password),
		)
		if err != nil {
			render.NotFoundf(w, "Invalid email or password")
			log.Debug().Err(err).
				Str("user", username).
				Msg("invalid password")
			return
		}

		expires := time.Now().Add(system.Config(ctx).Token.Expire)
		t, err := generate(user.ID, expires.Unix(), user.Salt)
		if err != nil {
			render.InternalErrorf(w, "Failed to create session")
			log.Debug().Err(err).
				Str("user", username).
				Msg("failed to generate token")
			return
		}

		token := &types.Token{
			Value:   t,
			Expires: expires.UTC(),
		}

		if r.FormValue("without_user") == "true" {
			render.JSON(w, token, 200)
		} else {
			render.JSON(w, &types.UserToken{
				User:  user,
				Token: token,
			}, 200)
		}

	}
}

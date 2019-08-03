// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/{{toLower repo}}/types/check"
	"github.com/rs/zerolog/hlog"
	"golang.org/x/crypto/bcrypt"

	"github.com/dchest/uniuri"
)

type userCreateInput struct {
	Username string `json:"email"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

// HandleCreate returns an http.HandlerFunc that processes an http.Request
// to create the named user account in the system.
func HandleCreate(users store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := hlog.FromRequest(r)

		in := new(userCreateInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			log.Debug().Err(err).
				Msg("cannot unmarshal json request")
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			render.InternalError(w, err)
			log.Debug().Err(err).
				Msg("cannot hash password")
			return
		}

		user := &types.User{
			Email:    in.Username,
			Admin:    in.Admin,
			Password: string(hash),
			Salt:     uniuri.NewLen(uniuri.UUIDLen),
			Created:  time.Now().Unix(),
			Updated:  time.Now().Unix(),
		}

		if ok, err := check.User(user); !ok {
			render.BadRequest(w, err)
			log.Debug().Err(err).
				Str("user_email", user.Email).
				Msg("cannot validate user")
			return
		}

		err = users.Create(ctx, user)
		if err != nil {
			render.InternalError(w, err)
			log.Error().Err(err).
				Int64("user_id", user.ID).
				Str("user_email", user.Email).
				Msg("cannot create user")
		} else {
			render.JSON(w, user, 200)
		}
	}
}

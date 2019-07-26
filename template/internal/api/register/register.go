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

package register

import (
	"net/http"
	"time"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/dchest/uniuri"
)

// HandleRegister returns an http.HandlerFunc that processes an http.Request
// to register the named user account with the system.
func HandleRegister(users store.UserStore, system store.SystemStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := r.FormValue("username")
		password := r.FormValue("password")

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot hash password")
			return
		}

		user := &types.User{
			Email:    username,
			Password: string(hash),
			Token:    uniuri.NewLen(uniuri.UUIDLen),
			Created:  time.Now().Unix(),
			Updated:  time.Now().Unix(),
		}

		if err := users.Create(ctx, user); err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("email", user.Email).
				Errorf("cannot create user")
			return
		}

		// if the registered user is the first user of the system,
		// assume they are the system administrator and grant the
		// user system admin access.
		if user.ID == 1 {
			user.Admin = true
			if err := users.Update(ctx, user); err != nil {
				logger.FromRequest(r).
					WithError(err).
					WithField("id", user.ID).
					WithField("email", user.Email).
					Errorf("cannot enable admin user")
			}
		}

		expires := time.Now().Add(system.Config(ctx).Token.Expire)
		token, err := generate(user.ID, expires.Unix(), user.Token)
		if err != nil {
			render.InternalErrorf(w, "Failed to create session")
			logger.FromRequest(r).
				WithError(err).
				WithField("user", user.Email).
				Errorln("failed to generate token")
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

// helper function generate a JWT token.
func generate(sub, exp int64, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": exp,
		"sub": sub,
		"iat": time.Now().Unix(),
	})
	return token.SignedString([]byte(secret))
}

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

package login

import (
	"net/http"
	"time"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HandleLogin returns an http.HandlerFunc that authenticates
// the user and returns an authentication token on success.
func HandleLogin(users store.UserStore, system store.SystemStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := r.FormValue("username")
		password := r.FormValue("password")
		user, err := users.FindEmail(ctx, username)
		if err != nil {
			render.NotFoundf(w, "Invalid email or password")
			logger.FromRequest(r).
				WithError(err).
				WithField("user", username).
				Debugln("cannot find user")
			return
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(user.Password),
			[]byte(password),
		)
		if err != nil {
			render.NotFoundf(w, "Invalid email or password")
			logger.FromRequest(r).
				WithError(err).
				WithField("user", username).
				Debugln("invalid password")
			return
		}

		expires := time.Now().Add(system.Config(ctx).Token.Expire)
		token, err := generate(user.ID, expires.Unix(), user.Token)
		if err != nil {
			render.InternalErrorf(w, "Failed to create session")
			logger.FromRequest(r).
				WithError(err).
				WithField("user", username).
				Debugln("failed to generate token")
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

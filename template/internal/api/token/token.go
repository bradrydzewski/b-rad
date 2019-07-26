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

package token

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/api/request"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"

	"github.com/dgrijalva/jwt-go"
)

// Must returns an http.HandlerFunc middleware that authenticates
// the http.Request and errors if the account cannot be authenticated.
func Must(users store.UserStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			str := extractToken(r)
			if len(str) == 0 {
				render.ErrorCode(w, errors.New("Requires authentication"), 401)
				return
			}

			var user *types.User
			token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
				var err error
				sub, ok := token.Claims.(jwt.MapClaims)["sub"].(float64)
				if !ok {
					return nil, errors.New("cannot read subscriber claim")
				}
				user, err = users.Find(ctx, int64(sub))
				if err != nil {
					logger.FromRequest(r).
						WithError(err).
						WithField("user", sub).
						Errorln("cannot find user")
					return nil, err
				}
				return []byte(user.Token), nil
			})
			if err != nil {
				render.ErrorCode(w, err, 401)
				return
			}
			if token.Valid == false {
				render.ErrorCode(w, errors.New("Invalid token"), 401)
				return
			}
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				render.ErrorCode(w, errors.New("Invalid token"), 401)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				if v, ok := claims["exp"]; ok {
					if time.Now().Unix() > int64(v.(float64)) {
						render.ErrorCode(w, errors.New("Expired token"), 401)
						return
					}
				}
			}

			log := logger.FromContext(ctx).
				WithField("user.email", user.Email).
				WithField("user.admin", user.Admin)

			ctx = logger.WithContext(ctx, log)
			next.ServeHTTP(w, r.WithContext(
				request.WithUser(ctx, user),
			))
		})
	}
}

func extractToken(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if bearer == "" {
		bearer = r.FormValue("access_token")
	}
	return strings.TrimPrefix(bearer, "Bearer ")
}

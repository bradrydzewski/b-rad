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

package user

import (
	"net/http"
	"time"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/api/request"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"

	"github.com/dgrijalva/jwt-go"
)

// HandleToken returns an http.HandlerFunc that generates and
// writes a json-encoded token to the http.Response body.
func HandleToken(users store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		token, err := generate(viewer.ID, viewer.Token)
		if err != nil {
			render.InternalErrorf(w, "Failed to generate token")
			logger.FromRequest(r).
				WithError(err).
				WithField("user", viewer.Email).
				Debugln("failed to generate token")
			return
		}

		render.JSON(w, &types.Token{Value: token}, 200)
	}
}

// helper function generate a JWT token.
var generate = func(sub int64, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Unix(),
	})
	return token.SignedString([]byte(secret))
}

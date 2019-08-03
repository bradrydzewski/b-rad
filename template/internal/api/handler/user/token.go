// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package user

import (
	"net/http"
	"time"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/api/request"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/rs/zerolog/hlog"

	"github.com/dgrijalva/jwt-go"
)

// HandleToken returns an http.HandlerFunc that generates and
// writes a json-encoded token to the http.Response body.
func HandleToken(users store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		token, err := generate(viewer.ID, viewer.Salt)
		if err != nil {
			render.InternalErrorf(w, "Failed to generate token")
			hlog.FromRequest(r).
				Error().Err(err).
				Str("user", viewer.Email).
				Msg("failed to generate token")
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

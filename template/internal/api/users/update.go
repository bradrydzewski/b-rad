// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package users

import (
	"encoding/json"
	"net/http"

	"github.com/{{toLower repo}}/internal/api/render"
	"github.com/{{toLower repo}}/internal/logger"
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"

	"github.com/go-chi/chi"
	"golang.org/x/crypto/bcrypt"
)

// GenerateFromPassword returns the bcrypt hash of the
// password at the given cost.
var hashPassword = bcrypt.GenerateFromPassword

// HandleUpdate returns an http.HandlerFunc that processes an http.Request
// to update a user account.
func HandleUpdate(users store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := new(types.UserInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Errorf("cannot unmarshal request")
			return
		}

		key := chi.URLParam(r, "user")
		user, err := users.FindKey(r.Context(), key)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("user", key).
				Errorf("cannot find user")
			return
		}

		if in.Password.IsZero() == false {
			hash, err := hashPassword([]byte(in.Password.String), bcrypt.DefaultCost)
			if err != nil {
				render.InternalError(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot hash password")
				return
			}
			user.Password = string(hash)
		}

		if in.Name.IsZero() == false {
			user.Name = in.Name.String
		}

		if in.Company.IsZero() == false {
			user.Company = in.Company.String
		}

		if in.Admin.Ptr() != nil {
			user.Admin = in.Admin.Bool
		}

		if in.Password.IsZero() == false {
			hash, err := bcrypt.GenerateFromPassword([]byte(in.Password.String), bcrypt.DefaultCost)
			if err != nil {
				render.InternalError(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot hash password")
				return
			}
			user.Password = string(hash)
		}

		err = users.Update(r.Context(), user)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", user.ID).
				WithField("email", user.Email).
				Errorf("cannot update user")
		} else {
			render.JSON(w, user, 200)
		}
	}
}

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

package users

import (
	"encoding/json"
	"net/http"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/chi"
	"gopkg.in/guregu/null.v4"
)

type userUpdateInput struct {
	Username null.String `json:"email"`
	Password null.String `json:"password"`
	Admin    null.Bool   `json:"admin"`
}

// HandleUpdate returns an http.HandlerFunc that processes an http.Request
// to update a user account.
func HandleUpdate(users store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := new(userUpdateInput)
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

		if in.Username.IsZero() == false {
			user.Email = in.Username.String
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

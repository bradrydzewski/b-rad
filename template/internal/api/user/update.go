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
	"encoding/json"
	"net/http"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/api/request"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"

	"golang.org/x/crypto/bcrypt"
)

// GenerateFromPassword returns the bcrypt hash of the
// password at the given cost.
var hashPassword = bcrypt.GenerateFromPassword

// HandleUpdate returns an http.HandlerFunc that processes an http.Request
// to update the current user account.
func HandleUpdate(users store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		in := new(types.UserInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("email", viewer.Email).
				Errorf("cannot unmarshal request")
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
			viewer.Password = string(hash)
		}

		if in.Username.IsZero() == false {
			viewer.Email = in.Username.String
		}

		err = users.Update(r.Context(), viewer)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("email", viewer.Email).
				Errorf("cannot update user")
		} else {
			render.JSON(w, viewer, 200)
		}
	}
}

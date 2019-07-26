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
	"time"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"
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

		in := new(userCreateInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot unmarshal json request")
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot hash password")
			return
		}

		user := &types.User{
			Email:    in.Username,
			Admin:    in.Admin,
			Password: string(hash),
			Token:    uniuri.NewLen(uniuri.UUIDLen),
			Created:  time.Now().Unix(),
			Updated:  time.Now().Unix(),
		}

		err = users.Create(r.Context(), user)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("email", user.Email).
				Errorf("cannot create user")
		} else {
			render.JSON(w, user, 200)
		}
	}
}

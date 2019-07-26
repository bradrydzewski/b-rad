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

// Package types defines common data structures.
package types

import (
	"time"

	"github.com/{{github}}/types/enum"
	"gopkg.in/guregu/null.v4"
)

type (
	// {{parent}} stores {{toLower parent}} details.
	{{parent}} struct {
		ID      int64  `db:"{{toLower parent}}_id"         json:"id"`
		Project int64  `db:"{{toLower parent}}_project_id" json:"-"`
		Name    string `db:"{{toLower parent}}_name"       json:"name"`
		Desc    string `db:"{{toLower parent}}_desc"       json:"desc"`
		Created int64  `db:"{{toLower parent}}_created"    json:"created"`
		Updated int64  `db:"{{toLower parent}}_updated"    json:"updated"`
	}

	// {{parent}}Input store details used to create or
	// update a {{toLower parent}}.
	{{parent}}Input struct {
		Name null.String `json:"name"`
		Desc null.String `json:"desc"`
	}

	// {{child}} stores {{toLower child}} details.
	{{child}} struct {
		ID      int64  `db:"{{toLower child}}_id"      json:"id"`
		{{parent}}     int64  `db:"{{toLower child}}_{{toLower parent}}_id"  json:"-"`
		Name    string `db:"{{toLower child}}_name"    json:"name"`
		Desc    string `db:"{{toLower child}}_desc"    json:"desc"`
		Created int64  `db:"{{toLower child}}_created" json:"created"`
		Updated int64  `db:"{{toLower child}}_updated" json:"updated"`
	}

	// {{child}}Input store details used to create or
	// update a {{toLower child}}.
	{{child}}Input struct {
		Name null.String `json:"name"`
		Desc null.String `json:"desc"`
	}

	// Member providers member details.
	Member struct {
		Email   string    `db:"user_email"        json:"email"`
		Project int64     `db:"member_project_id" json:"-"`
		User    int64     `db:"member_user_id"    json:"-"`
		Role    enum.Role `db:"member_role"       json:"role"`
	}

	// Membership stores membership details.
	Membership struct {
		Project int64     `db:"member_project_id" json:"-"`
		User    int64     `db:"member_user_id"    json:"-"`
		Role    enum.Role `db:"member_role"       json:"role"`
	}

	// MembershipInput stores membership details.
	MembershipInput struct {
		Project int64     `db:"member_project_id" json:"project"`
		User    string    `db:"member_user_id"    json:"user"`
		Role    enum.Role `db:"member_role"       json:"role"`
	}

	// Params stores query parameters.
	Params struct {
		Page int `json:"page"`
		Size int `json:"size"`
	}

	// Project stores project details.
	Project struct {
		ID      int64  `db:"project_id"      json:"id"`
		Name    string `db:"project_name"    json:"name"`
		Desc    string `db:"project_desc"    json:"desc"`
		Token   string `db:"project_token"   json:"-"`
		Active  bool   `db:"project_active"  json:"active"`
		Created int64  `db:"project_created" json:"created"`
		Updated int64  `db:"project_updated" json:"updated"`
	}

	// ProjectInput store user project details used to
	// create or update a project.
	ProjectInput struct {
		Name null.String `json:"name"`
		Desc null.String `json:"desc"`
	}

	// Token stores token  details.
	Token struct {
		Value   string    `json:"access_token"`
		Address string    `json:"uri,omitempty"`
		Expires time.Time `json:"expires_at,omitempty"`
	}

	// User stores user account details.
	User struct {
		ID       int64  `db:"user_id"        json:"id"`
		Email    string `db:"user_email"     json:"email"`
		Password string `db:"user_password"  json:"-"`
		Token    string `db:"user_token"     json:"-"`
		Admin    bool   `db:"user_admin"     json:"admin"`
		Blocked  bool   `db:"user_blocked"   json:"-"`
		Created  int64  `db:"user_created"   json:"created"`
		Updated  int64  `db:"user_updated"   json:"updated"`
		Authed   int64  `db:"user_authed"    json:"authed"`
	}

	// UserInput store user account details used to
	// create or update a user.
	UserInput struct {
		Username null.String `json:"email"`
		Password null.String `json:"password"`
		Admin    null.Bool   `json:"admin"`
	}

	// UserToken stores user account and token details.
	UserToken struct {
		User  *User  `json:"user"`
		Token *Token `json:"token"`
	}
)

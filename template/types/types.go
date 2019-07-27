// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package types defines common data structures.
package types

import (
	"time"

	"github.com/{{toLower repo}}/types/enum"
	"gopkg.in/guregu/null.v4"
)

type (
	// {{title parent}} stores {{toLower parent}} details.
	{{title parent}} struct {
		ID      int64  `db:"{{toLower parent}}_id"         json:"id"`
		{{title project}} int64  `db:"{{toLower parent}}_{{toLower project}}_id" json:"-"`
		Name    string `db:"{{toLower parent}}_name"       json:"name"`
		Desc    string `db:"{{toLower parent}}_desc"       json:"desc"`
		Created int64  `db:"{{toLower parent}}_created"    json:"created"`
		Updated int64  `db:"{{toLower parent}}_updated"    json:"updated"`
	}

	// {{title parent}}Input store details used to create or
	// update a {{toLower parent}}.
	{{title parent}}Input struct {
		Name null.String `json:"name"`
		Desc null.String `json:"desc"`
	}

	// {{title child}} stores {{toLower child}} details.
	{{title child}} struct {
		ID      int64  `db:"{{toLower child}}_id"      json:"id"`
		{{title parent}}   int64  `db:"{{toLower child}}_{{toLower parent}}_id"  json:"-"`
		Name    string `db:"{{toLower child}}_name"    json:"name"`
		Desc    string `db:"{{toLower child}}_desc"    json:"desc"`
		Created int64  `db:"{{toLower child}}_created" json:"created"`
		Updated int64  `db:"{{toLower child}}_updated" json:"updated"`
	}

	// {{title child}}Input store details used to create or
	// update a {{toLower child}}.
	{{title child}}Input struct {
		Name null.String `json:"name"`
		Desc null.String `json:"desc"`
	}

	// Member providers member details.
	Member struct {
		Email   string    `db:"user_email"        json:"email"`
		{{title project}} int64     `db:"member_{{toLower project}}_id" json:"-"`
		User    int64     `db:"member_user_id"    json:"-"`
		Role    enum.Role `db:"member_role"       json:"role"`
	}

	// Membership stores membership details.
	Membership struct {
		{{title project}} int64     `db:"member_{{toLower project}}_id" json:"-"`
		User    int64     `db:"member_user_id"    json:"-"`
		Role    enum.Role `db:"member_role"       json:"role"`
	}

	// MembershipInput stores membership details.
	MembershipInput struct {
		{{title project}} int64     `db:"member_{{toLower project}}_id" json:"{{toLower project}}"`
		User    string    `db:"member_user_id"    json:"user"`
		Role    enum.Role `db:"member_role"       json:"role"`
	}

	// Params stores query parameters.
	Params struct {
		Page int `json:"page"`
		Size int `json:"size"`
	}

	// {{title project}} stores {{toLower project}} details.
	{{title project}} struct {
		ID      int64  `db:"{{toLower project}}_id"      json:"id"`
		Name    string `db:"{{toLower project}}_name"    json:"name"`
		Desc    string `db:"{{toLower project}}_desc"    json:"desc"`
		Token   string `db:"{{toLower project}}_token"   json:"-"`
		Active  bool   `db:"{{toLower project}}_active"  json:"active"`
		Created int64  `db:"{{toLower project}}_created" json:"created"`
		Updated int64  `db:"{{toLower project}}_updated" json:"updated"`
	}

	// {{title project}}Input store user {{toLower project}} details used to
	// create or update a {{toLower project}}.
	{{title project}}Input struct {
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
		Name     string `db:"user_name"      json:"name"`
		Company  string `db:"user_company"   json:"company"`
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
		Name     null.String `json:"name"`
		Company  null.String `json:"company"`
		Admin    null.Bool   `json:"admin"`
	}

	// UserToken stores user account and token details.
	UserToken struct {
		User  *User  `json:"user"`
		Token *Token `json:"token"`
	}
)

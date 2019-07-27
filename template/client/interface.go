// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package client

import "github.com/{{toLower repo}}/types"

// Client to access the remote APIs.
type Client interface {
	// Login authenticates the user and returns a JWT token.
	Login(username, password string) (*types.Token, error)

	// Register registers a new  user and returns a JWT token.
	Register(username, password string) (*types.Token, error)

	// Self returns the currently authenticated user.
	Self() (*types.User, error)

	// Token returns an oauth2 bearer token for the currently
	// authenticated user.
	Token() (*types.Token, error)

	// User returns a user by ID or email.
	User(key string) (*types.User, error)

	// UserList returns a list of all registered users.
	UserList() ([]*types.User, error)

	// UserCreate creates a new user account.
	UserCreate(user *types.User) (*types.User, error)

	// UserUpdate updates a user account by ID or email.
	UserUpdate(key string, input *types.UserInput) (*types.User, error)

	// UserDelete deletes a user account by ID or email.
	UserDelete(key string) error

	// {{title project}} returns a {{toLower project}} by ID.
	{{title project}}(id int64) (*types.{{title project}}, error)

	// {{title project}}List returns a list of all {{toLower project}}s.
	{{title project}}List() ([]*types.{{title project}}, error)

	// {{title project}}Create creates a new {{toLower project}}.
	{{title project}}Create(user *types.{{title project}}) (*types.{{title project}}, error)

	// {{title project}}Update updates a {{toLower project}}.
	{{title project}}Update(id int64, input *types.{{title project}}Input) (*types.{{title project}}, error)

	// {{title project}}Delete deletes a {{toLower project}}.
	{{title project}}Delete(id int64) error

	// Member returns a membrer by ID.
	Member({{toLower project}} int64, user string) (*types.Member, error)

	// MemberList returns a list of all {{toLower project}} members.
	MemberList({{toLower project}} int64) ([]*types.Member, error)

	// MemberCreate creates a new {{toLower project}} member.
	MemberCreate(member *types.MembershipInput) (*types.Member, error)

	// MemberUpdate updates a {{toLower project}} member.
	MemberUpdate(member *types.MembershipInput) (*types.Member, error)

	// MemberDelete deletes a {{toLower project}} member.
	MemberDelete({{toLower project}} int64, user string) error

	// {{title parent}} returns a {{toLower parent}} by ID.
	{{title parent}}({{toLower project}}, id int64) (*types.{{title parent}}, error)

	// {{title parent}}List returns a list of all {{toLower parent}}s by {{toLower project}} id.
	{{title parent}}List({{toLower project}} int64) ([]*types.{{title parent}}, error)

	// {{title parent}}Create creates a new {{toLower parent}}.
	{{title parent}}Create({{toLower project}} int64, {{toLower parent}} *types.{{title parent}}) (*types.{{title parent}}, error)

	// {{title parent}}Update updates a {{toLower parent}}.
	{{title parent}}Update({{toLower project}}, id int64, input *types.{{title parent}}Input) (*types.{{title parent}}, error)

	// {{title parent}}Delete deletes a {{toLower parent}}.
	{{title parent}}Delete({{toLower project}}, id int64) error

	// {{title child}} returns a {{toLower child}} by ID.
	{{title child}}({{toLower project}}, {{toLower parent}}, {{toLower child}} int64) (*types.{{title child}}, error)

	// {{title child}}List returns a list of all {{toLower child}}s by {{toLower project}} id.
	{{title child}}List({{toLower project}}, {{toLower parent}} int64) ([]*types.{{title child}}, error)

	// {{title child}}Create creates a new {{toLower child}}.
	{{title child}}Create({{toLower project}}, {{toLower parent}} int64, input *types.{{title child}}) (*types.{{title child}}, error)

	// {{title child}}Update updates a {{toLower child}}.
	{{title child}}Update({{toLower project}}, {{toLower parent}}, {{toLower child}} int64, input *types.{{title child}}Input) (*types.{{title child}}, error)

	// {{title child}}Delete deletes a {{toLower child}}.
	{{title child}}Delete({{toLower project}}, {{toLower parent}}, {{toLower child}} int64) error
}

// remoteError store the error payload returned
// fro the remote API.
type remoteError struct {
	Message string `json:"message"`
}

// Error returns the error message.
func (e *remoteError) Error() string {
	return e.Message
}

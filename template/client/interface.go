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

package client

import "github.com/{{github}}/types"

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

	// Project returns a project by ID.
	Project(id int64) (*types.Project, error)

	// ProjectList returns a list of all projects.
	ProjectList() ([]*types.Project, error)

	// ProjectCreate creates a new project.
	ProjectCreate(user *types.Project) (*types.Project, error)

	// ProjectUpdate updates a project.
	ProjectUpdate(id int64, input *types.ProjectInput) (*types.Project, error)

	// ProjectDelete deletes a project.
	ProjectDelete(id int64) error

	// Member returns a membrer by ID.
	Member(project int64, user string) (*types.Member, error)

	// MemberList returns a list of all project members.
	MemberList(project int64) ([]*types.Member, error)

	// MemberCreate creates a new project member.
	MemberCreate(member *types.MembershipInput) (*types.Member, error)

	// MemberUpdate updates a project member.
	MemberUpdate(member *types.MembershipInput) (*types.Member, error)

	// MemberDelete deletes a project member.
	MemberDelete(project int64, user string) error

	// {{parent}} returns a {{toLower parent}} by ID.
	{{parent}}(project, id int64) (*types.{{parent}}, error)

	// {{parent}}List returns a list of all {{toLower parent}}s by project id.
	{{parent}}List(project int64) ([]*types.{{parent}}, error)

	// {{parent}}Create creates a new {{toLower parent}}.
	{{parent}}Create(project int64, {{toLower parent}} *types.{{parent}}) (*types.{{parent}}, error)

	// {{parent}}Update updates a {{toLower parent}}.
	{{parent}}Update(project, id int64, input *types.{{parent}}Input) (*types.{{parent}}, error)

	// {{parent}}Delete deletes a {{toLower parent}}.
	{{parent}}Delete(project, id int64) error

	// {{child}} returns a {{toLower child}} by ID.
	{{child}}(project, {{toLower parent}}, {{toLower child}} int64) (*types.{{child}}, error)

	// {{child}}List returns a list of all {{toLower child}}s by project id.
	{{child}}List(project, {{toLower parent}} int64) ([]*types.{{child}}, error)

	// {{child}}Create creates a new {{toLower child}}.
	{{child}}Create(project, {{toLower parent}} int64, input *types.{{child}}) (*types.{{child}}, error)

	// {{child}}Update updates a {{toLower child}}.
	{{child}}Update(project, {{toLower parent}}, {{toLower child}} int64, input *types.{{child}}Input) (*types.{{child}}, error)

	// {{child}}Delete deletes a {{toLower child}}.
	{{child}}Delete(project, {{toLower parent}}, {{toLower child}} int64) error
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

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

// Package store defines the data storage interfaces.
package store

import (
	"context"

	"github.com/{{github}}/types"
)

type (
	// {{child}}Store defines {{toLower child}} data storage.
	{{child}}Store interface {
		// Find finds the {{toLower child}} by id.
		Find(ctx context.Context, id int64) (*types.{{child}}, error)

		// List returns a list of {{toLower child}}s by {{toLower parent}} id.
		List(ctx context.Context, id int64, params types.Params) ([]*types.{{child}}, error)

		// Create saves the {{toLower child}} details.
		Create(ctx context.Context, {{toLower child}} *types.{{child}}) error

		// Update updates the {{toLower child}} details.
		Update(ctx context.Context, {{toLower child}} *types.{{child}}) error

		// Delete deletes the {{toLower child}}.
		Delete(ctx context.Context, {{toLower child}} *types.{{child}}) error
	}

	// {{parent}}Store defines {{toLower parent}} data storage.
	{{parent}}Store interface {
		// Find finds the {{toLower parent}} by id.
		Find(ctx context.Context, id int64) (*types.{{parent}}, error)

		// List returns a list of {{toLower parent}}s by account id.
		List(ctx context.Context, id int64, params types.Params) ([]*types.{{parent}}, error)

		// Create saves the {{toLower parent}} details.
		Create(ctx context.Context, {{toLower parent}} *types.{{parent}}) error

		// Update updates the {{toLower parent}} details.
		Update(ctx context.Context, {{toLower parent}} *types.{{parent}}) error

		// Delete deletes the {{toLower parent}}.
		Delete(ctx context.Context, {{toLower parent}} *types.{{parent}}) error
	}

	// MemberStore defines member data storage.
	MemberStore interface {
		// Find finds the member by project and user id.
		Find(ctx context.Context, project, user int64) (*types.Member, error)

		// List returns a list of members.
		List(ctx context.Context, project int64, params types.Params) ([]*types.Member, error)

		// Create saves the membership details.
		Create(ctx context.Context, membership *types.Membership) error

		// Update updates the membership details.
		Update(ctx context.Context, membership *types.Membership) error

		// Delete deletes the membership.
		Delete(ctx context.Context, project, user int64) error
	}

	// ProjectStore defines project data storage.
	ProjectStore interface {
		// Find finds the project by id.
		Find(ctx context.Context, id int64) (*types.Project, error)

		// FindToken finds the project by token.
		FindToken(ctx context.Context, token string) (*types.Project, error)

		// List returns a list of projects by user.
		List(ctx context.Context, user int64, params types.Params) ([]*types.Project, error)

		// Create saves the project details.
		Create(ctx context.Context, project *types.Project) error

		// Update updates the project details.
		Update(ctx context.Context, project *types.Project) error

		// Delete deletes the project.
		Delete(ctx context.Context, project *types.Project) error
	}

	// UserStore defines user data storage.
	UserStore interface {
		// Find finds the user by id.
		Find(ctx context.Context, id int64) (*types.User, error)

		// FindEmail finds the user by email.
		FindEmail(ctx context.Context, email string) (*types.User, error)

		// FindKey finds the user by unique key (email or id).
		FindKey(ctx context.Context, key string) (*types.User, error)

		// List returns a list of users.
		List(ctx context.Context, params types.Params) ([]*types.User, error)

		// Create saves the user details.
		Create(ctx context.Context, user *types.User) error

		// Update updates the user details.
		Update(ctx context.Context, user *types.User) error

		// Delete deletes the user.
		Delete(ctx context.Context, user *types.User) error

		// Count returns a count of users.
		Count(ctx context.Context) (int64, error)
	}

	// SystemStore defines insternal system metadata storage.
	SystemStore interface {
		// Config returns the system configuration.
		Config(ctx context.Context) *types.Config
	}
)

// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package store defines the data storage interfaces.
package store

import (
	"context"

	"github.com/{{toLower repo}}/types"
)

type (
	// {{title child}}Store defines {{toLower child}} data storage.
	{{title child}}Store interface {
		// Find finds the {{toLower child}} by id.
		Find(ctx context.Context, id int64) (*types.{{title child}}, error)

		// List returns a list of {{toLower child}}s by {{toLower parent}} id.
		List(ctx context.Context, id int64, params types.Params) ([]*types.{{title child}}, error)

		// Create saves the {{toLower child}} details.
		Create(ctx context.Context, {{toLower child}} *types.{{title child}}) error

		// Update updates the {{toLower child}} details.
		Update(ctx context.Context, {{toLower child}} *types.{{title child}}) error

		// Delete deletes the {{toLower child}}.
		Delete(ctx context.Context, {{toLower child}} *types.{{title child}}) error
	}

	// {{title parent}}Store defines {{toLower parent}} data storage.
	{{title parent}}Store interface {
		// Find finds the {{toLower parent}} by id.
		Find(ctx context.Context, id int64) (*types.{{title parent}}, error)

		// List returns a list of {{toLower parent}}s by account id.
		List(ctx context.Context, id int64, params types.Params) ([]*types.{{title parent}}, error)

		// Create saves the {{toLower parent}} details.
		Create(ctx context.Context, {{toLower parent}} *types.{{title parent}}) error

		// Update updates the {{toLower parent}} details.
		Update(ctx context.Context, {{toLower parent}} *types.{{title parent}}) error

		// Delete deletes the {{toLower parent}}.
		Delete(ctx context.Context, {{toLower parent}} *types.{{title parent}}) error
	}

	// MemberStore defines member data storage.
	MemberStore interface {
		// Find finds the member by {{toLower project}} and user id.
		Find(ctx context.Context, {{toLower project}}, user int64) (*types.Member, error)

		// List returns a list of members.
		List(ctx context.Context, {{toLower project}} int64, params types.Params) ([]*types.Member, error)

		// Create saves the membership details.
		Create(ctx context.Context, membership *types.Membership) error

		// Update updates the membership details.
		Update(ctx context.Context, membership *types.Membership) error

		// Delete deletes the membership.
		Delete(ctx context.Context, {{toLower project}}, user int64) error
	}

	// {{title project}}Store defines {{toLower project}} data storage.
	{{title project}}Store interface {
		// Find finds the {{toLower project}} by id.
		Find(ctx context.Context, id int64) (*types.{{title project}}, error)

		// FindToken finds the {{toLower project}} by token.
		FindToken(ctx context.Context, token string) (*types.{{title project}}, error)

		// List returns a list of {{toLower project}}s by user.
		List(ctx context.Context, user int64, params types.Params) ([]*types.{{title project}}, error)

		// Create saves the {{toLower project}} details.
		Create(ctx context.Context, {{toLower project}} *types.{{title project}}) error

		// Update updates the {{toLower project}} details.
		Update(ctx context.Context, {{toLower project}} *types.{{title project}}) error

		// Delete deletes the {{toLower project}}.
		Delete(ctx context.Context, {{toLower project}} *types.{{title project}}) error
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

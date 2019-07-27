// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// WireSet provides a wire set for this package
var WireSet = wire.NewSet(
	ProvideDatabase,
	ProvideUserStore,
	Provide{{title project}}Store,
	ProvideMemberStore,
	Provide{{title parent}}Store,
	Provide{{title child}}Store,
)

// ProvideDatabase provides a database connection.
func ProvideDatabase(config *types.Config) (*sqlx.DB, error) {
	return Connect(
		config.Database.Driver,
		config.Database.Datasource,
	)
}

// ProvideUserStore provides a user store.
func ProvideUserStore(db *sqlx.DB) store.UserStore {
	switch db.DriverName() {
	case "postgres":
		return NewUserStore(db)
	default:
		return NewUserStoreSync(
			NewUserStore(db),
		)
	}
}

// Provide{{title project}}Store provides a {{toLower project}} store.
func Provide{{title project}}Store(db *sqlx.DB) store.{{title project}}Store {
	switch db.DriverName() {
	case "postgres":
		return New{{title project}}Store(db)
	default:
		return New{{title project}}StoreSync(
			New{{title project}}Store(db),
		)
	}
}

// ProvideMemberStore provides a member store.
func ProvideMemberStore(db *sqlx.DB) store.MemberStore {
	switch db.DriverName() {
	case "postgres":
		return NewMemberStore(db)
	default:
		return NewMemberStoreSync(
			NewMemberStore(db),
		)
	}
}

// Provide{{title parent}}Store provides a {{toLower parent}} store.
func Provide{{title parent}}Store(db *sqlx.DB) store.{{title parent}}Store {
	switch db.DriverName() {
	case "postgres":
		return New{{title parent}}Store(db)
	default:
		return New{{title parent}}StoreSync(
			New{{title parent}}Store(db),
		)
	}
}

// Provide{{title child}}Store provides a {{toLower child}} store.
func Provide{{title child}}Store(db *sqlx.DB) store.{{title child}}Store {
	switch db.DriverName() {
	case "postgres":
		return New{{title child}}Store(db)
	default:
		return New{{title child}}StoreSync(
			New{{title child}}Store(db),
		)
	}
}

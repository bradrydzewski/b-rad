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

package database

import (
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// WireSet provides a wire set for this package
var WireSet = wire.NewSet(
	ProvideDatabase,
	ProvideUserStore,
	ProvideProjectStore,
	ProvideMemberStore,
	Provide{{parent}}Store,
	Provide{{child}}Store,
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

// ProvideProjectStore provides a project store.
func ProvideProjectStore(db *sqlx.DB) store.ProjectStore {
	switch db.DriverName() {
	case "postgres":
		return NewProjectStore(db)
	default:
		return NewProjectStoreSync(
			NewProjectStore(db),
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

// Provide{{parent}}Store provides a {{toLower parent}} store.
func Provide{{parent}}Store(db *sqlx.DB) store.{{parent}}Store {
	switch db.DriverName() {
	case "postgres":
		return New{{parent}}Store(db)
	default:
		return New{{parent}}StoreSync(
			New{{parent}}Store(db),
		)
	}
}

// Provide{{child}}Store provides a {{toLower child}} store.
func Provide{{child}}Store(db *sqlx.DB) store.{{child}}Store {
	switch db.DriverName() {
	case "postgres":
		return New{{child}}Store(db)
	default:
		return New{{child}}StoreSync(
			New{{child}}Store(db),
		)
	}
}

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
	"context"

	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/internal/store/database/mutex"
	"github.com/{{github}}/types"
)

var _ store.{{child}}Store = (*{{child}}StoreSync)(nil)

// New{{child}}StoreSync returns a new {{child}}StoreSync.
func New{{child}}StoreSync(store *{{child}}Store) *{{child}}StoreSync {
	return &{{child}}StoreSync{store}
}

// {{child}}StoreSync synronizes read and write access to the
// {{toLower child}} store. This prevents race conditions when the database
// type is sqlite3.
type {{child}}StoreSync struct{ *{{child}}Store }

// Find finds the {{toLower child}} by id.
func (s *{{child}}StoreSync) Find(ctx context.Context, id int64) (*types.{{child}}, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.{{child}}Store.Find(ctx, id)
}

// List returns a list of {{toLower child}}s.
func (s *{{child}}StoreSync) List(ctx context.Context, id int64, opts types.Params) ([]*types.{{child}}, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.{{child}}Store.List(ctx, id, opts)
}

// Create saves the {{toLower child}} details.
func (s *{{child}}StoreSync) Create(ctx context.Context, {{toLower child}} *types.{{child}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{child}}Store.Create(ctx, {{toLower child}})
}

// Update updates the {{toLower child}} details.
func (s *{{child}}StoreSync) Update(ctx context.Context, {{toLower child}} *types.{{child}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{child}}Store.Update(ctx, {{toLower child}})
}

// Delete deletes the {{toLower child}}.
func (s *{{child}}StoreSync) Delete(ctx context.Context, {{toLower child}} *types.{{child}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{child}}Store.Delete(ctx, {{toLower child}})
}

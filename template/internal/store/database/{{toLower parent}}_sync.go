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

var _ store.{{parent}}Store = (*{{parent}}StoreSync)(nil)

// New{{parent}}StoreSync returns a new {{parent}}StoreSync.
func New{{parent}}StoreSync(store *{{parent}}Store) *{{parent}}StoreSync {
	return &{{parent}}StoreSync{store}
}

// {{parent}}StoreSync synronizes read and write access to the
// {{toLower parent}} store. This prevents race conditions when the database
// type is sqlite3.
type {{parent}}StoreSync struct{ *{{parent}}Store }

// Find finds the {{toLower parent}} by id.
func (s *{{parent}}StoreSync) Find(ctx context.Context, id int64) (*types.{{parent}}, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.{{parent}}Store.Find(ctx, id)
}

// List returns a list of {{toLower parent}}s.
func (s *{{parent}}StoreSync) List(ctx context.Context, id int64, opts types.Params) ([]*types.{{parent}}, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.{{parent}}Store.List(ctx, id, opts)
}

// Create saves the {{toLower parent}} details.
func (s *{{parent}}StoreSync) Create(ctx context.Context, {{toLower parent}} *types.{{parent}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{parent}}Store.Create(ctx, {{toLower parent}})
}

// Update updates the {{toLower parent}} details.
func (s *{{parent}}StoreSync) Update(ctx context.Context, {{toLower parent}} *types.{{parent}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{parent}}Store.Update(ctx, {{toLower parent}})
}

// Delete deletes the {{toLower parent}}.
func (s *{{parent}}StoreSync) Delete(ctx context.Context, {{toLower parent}} *types.{{parent}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{parent}}Store.Delete(ctx, {{toLower parent}})
}

// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"context"

	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/internal/store/database/mutex"
	"github.com/{{toLower repo}}/types"
)

var _ store.{{title child}}Store = (*{{title child}}StoreSync)(nil)

// New{{title child}}StoreSync returns a new {{title child}}StoreSync.
func New{{title child}}StoreSync(store *{{title child}}Store) *{{title child}}StoreSync {
	return &{{title child}}StoreSync{store}
}

// {{title child}}StoreSync synronizes read and write access to the
// {{toLower child}} store. This prevents race conditions when the database
// type is sqlite3.
type {{title child}}StoreSync struct{ *{{title child}}Store }

// Find finds the {{toLower child}} by id.
func (s *{{title child}}StoreSync) Find(ctx context.Context, id int64) (*types.{{title child}}, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.{{title child}}Store.Find(ctx, id)
}

// List returns a list of {{toLower child}}s.
func (s *{{title child}}StoreSync) List(ctx context.Context, id int64, opts types.Params) ([]*types.{{title child}}, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.{{title child}}Store.List(ctx, id, opts)
}

// Create saves the {{toLower child}} details.
func (s *{{title child}}StoreSync) Create(ctx context.Context, {{toLower child}} *types.{{title child}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{title child}}Store.Create(ctx, {{toLower child}})
}

// Update updates the {{toLower child}} details.
func (s *{{title child}}StoreSync) Update(ctx context.Context, {{toLower child}} *types.{{title child}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{title child}}Store.Update(ctx, {{toLower child}})
}

// Delete deletes the {{toLower child}}.
func (s *{{title child}}StoreSync) Delete(ctx context.Context, {{toLower child}} *types.{{title child}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{title child}}Store.Delete(ctx, {{toLower child}})
}

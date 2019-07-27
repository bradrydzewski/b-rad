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

var _ store.{{title project}}Store = (*{{title project}}StoreSync)(nil)

// New{{title project}}StoreSync returns a new {{title project}}StoreSync.
func New{{title project}}StoreSync(store *{{title project}}Store) *{{title project}}StoreSync {
	return &{{title project}}StoreSync{store}
}

// {{title project}}StoreSync synronizes read and write access to the
// {{toLower project}} store. This prevents race conditions when the database
// type is sqlite3.
type {{title project}}StoreSync struct{ *{{title project}}Store }

// Find finds the {{toLower project}} by id.
func (s *{{title project}}StoreSync) Find(ctx context.Context, id int64) (*types.{{title project}}, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.{{title project}}Store.Find(ctx, id)
}

// List returns a list of {{toLower project}}s by user.
func (s *{{title project}}StoreSync) List(ctx context.Context, id int64, opts types.Params) ([]*types.{{title project}}, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.{{title project}}Store.List(ctx, id, opts)
}

// Create saves the {{toLower project}} details.
func (s *{{title project}}StoreSync) Create(ctx context.Context, {{toLower project}} *types.{{title project}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{title project}}Store.Create(ctx, {{toLower project}})
}

// Update updates the {{toLower project}} details.
func (s *{{title project}}StoreSync) Update(ctx context.Context, {{toLower project}} *types.{{title project}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{title project}}Store.Update(ctx, {{toLower project}})
}

// Delete deletes the {{toLower project}}.
func (s *{{title project}}StoreSync) Delete(ctx context.Context, {{toLower project}} *types.{{title project}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{title project}}Store.Delete(ctx, {{toLower project}})
}

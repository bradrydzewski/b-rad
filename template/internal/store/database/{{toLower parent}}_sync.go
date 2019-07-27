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

var _ store.{{title parent}}Store = (*{{title parent}}StoreSync)(nil)

// New{{title parent}}StoreSync returns a new {{title parent}}StoreSync.
func New{{title parent}}StoreSync(store *{{title parent}}Store) *{{title parent}}StoreSync {
	return &{{title parent}}StoreSync{store}
}

// {{title parent}}StoreSync synronizes read and write access to the
// {{toLower parent}} store. This prevents race conditions when the database
// type is sqlite3.
type {{title parent}}StoreSync struct{ *{{title parent}}Store }

// Find finds the {{toLower parent}} by id.
func (s *{{title parent}}StoreSync) Find(ctx context.Context, id int64) (*types.{{title parent}}, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.{{title parent}}Store.Find(ctx, id)
}

// List returns a list of {{toLower parent}}s.
func (s *{{title parent}}StoreSync) List(ctx context.Context, id int64, opts types.Params) ([]*types.{{title parent}}, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.{{title parent}}Store.List(ctx, id, opts)
}

// Create saves the {{toLower parent}} details.
func (s *{{title parent}}StoreSync) Create(ctx context.Context, {{toLower parent}} *types.{{title parent}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{title parent}}Store.Create(ctx, {{toLower parent}})
}

// Update updates the {{toLower parent}} details.
func (s *{{title parent}}StoreSync) Update(ctx context.Context, {{toLower parent}} *types.{{title parent}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{title parent}}Store.Update(ctx, {{toLower parent}})
}

// Delete deletes the {{toLower parent}}.
func (s *{{title parent}}StoreSync) Delete(ctx context.Context, {{toLower parent}} *types.{{title parent}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{title parent}}Store.Delete(ctx, {{toLower parent}})
}

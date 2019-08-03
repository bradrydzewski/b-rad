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

var _ store.MemberStore = (*MemberStoreSync)(nil)

// NewMemberStoreSync returns a new MemberStoreSync.
func NewMemberStoreSync(store *MemberStore) *MemberStoreSync {
	return &MemberStoreSync{base: store}
}

// MemberStoreSync synronizes read and write access to the
// membership store. This prevents race conditions when the database
// type is sqlite3.
type MemberStoreSync struct{ base *MemberStore }

// Find finds the member by {{toLower project}} and user id.
func (s *MemberStoreSync) Find(ctx context.Context, {{toLower project}} int64, user int64) (*types.Member, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.base.Find(ctx, {{toLower project}}, user)
}

// List returns a list of members.
func (s *MemberStoreSync) List(ctx context.Context, {{toLower project}} int64, opts types.Params) ([]*types.Member, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.base.List(ctx, {{toLower project}}, opts)
}

// Create saves the membership details.
func (s *MemberStoreSync) Create(ctx context.Context, membership *types.Membership) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.base.Create(ctx, membership)
}

// Update updates the membership details.
func (s *MemberStoreSync) Update(ctx context.Context, membership *types.Membership) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.base.Update(ctx, membership)
}

// Delete deletes the membership.
func (s *MemberStoreSync) Delete(ctx context.Context, {{toLower project}}, user int64) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.base.Delete(ctx, {{toLower project}}, user)
}

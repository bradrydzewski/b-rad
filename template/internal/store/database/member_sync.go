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

var _ store.MemberStore = (*MemberStoreSync)(nil)

// NewMemberStoreSync returns a new MemberStoreSync.
func NewMemberStoreSync(store *MemberStore) *MemberStoreSync {
	return &MemberStoreSync{store}
}

// MemberStoreSync synronizes read and write access to the
// membership store. This prevents race conditions when the database
// type is sqlite3.
type MemberStoreSync struct{ *MemberStore }

// Find finds the member by project and user id.
func (s *MemberStoreSync) Find(ctx context.Context, project int64, user int64) (*types.Member, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.MemberStore.Find(ctx, project, user)
}

// List returns a list of members.
func (s *MemberStoreSync) List(ctx context.Context, project int64, opts types.Params) ([]*types.Member, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.MemberStore.List(ctx, project, opts)
}

// Create saves the membership details.
func (s *MemberStoreSync) Create(ctx context.Context, membership *types.Membership) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.MemberStore.Create(ctx, membership)
}

// Update updates the membership details.
func (s *MemberStoreSync) Update(ctx context.Context, membership *types.Membership) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.MemberStore.Update(ctx, membership)
}

// Delete deletes the membership.
func (s *MemberStoreSync) Delete(ctx context.Context, project, user int64) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.MemberStore.Delete(ctx, project, user)
}

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
	"database/sql"
	"testing"

	"github.com/{{github}}/types"
	"github.com/{{github}}/types/enum"
)

func TestMembershipFind(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err := seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewMemberStoreSync(&MemberStore{db})
	result, err := store.Find(noContext, 1, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if got, want := result.Project, int64(1); want != got {
		t.Errorf("Want account ID %d, got %d", want, got)
	}
	if got, want := result.User, int64(2); want != got {
		t.Errorf("Want user ID %d, got %d", want, got)
	}
	if got, want := result.Email, "john@example.com"; want != got {
		t.Errorf("Want user email %q, got %q", want, got)
	}
}

func TestMembershipList(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err := seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewMemberStoreSync(NewMemberStore(db))
	result, err := store.List(noContext, 1, types.Params{})
	if err != nil {
		t.Error(err)
		return
	}
	if got, want := len(result), 2; want != got {
		t.Errorf("Want member count %d, got %d", want, got)
	}
	if got, want := result[0].Project, int64(1); want != got {
		t.Errorf("Want account ID %d, got %d", want, got)
	}
	if got, want := result[0].User, int64(1); want != got {
		t.Errorf("Want user ID %d, got %d", want, got)
	}
	if got, want := result[0].Email, "jane@example.com"; want != got {
		t.Errorf("Want user email %q, got %q", want, got)
	}
}

func TestMembershipCreate(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err := seed(db); err != nil {
		t.Error(err)
		return
	}
	store := NewMemberStoreSync(NewMemberStore(db))
	create := &types.Membership{
		Project: 2,
		User:    2,
		Role:    enum.RoleAdmin,
	}
	if err := store.Create(noContext, create); err != nil {
		t.Error(err)
		return
	}

	found, err := store.Find(noContext, 2, 2)
	if err != nil {
		t.Error(err)
		return
	}

	if got, want := found.Role, found.Role; got != want {
		t.Errorf("Want role %v, got %v", want, got)
	}
}

func TestMembershipUpdate(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err := seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewMemberStoreSync(NewMemberStore(db))
	result, err := store.Find(noContext, 1, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if result.Role != enum.RoleAdmin {
		t.Errorf("Expected admin role, got role %s", result.Role)
		return
	}

	result.Role = enum.RoleDeveloper
	err = store.Update(noContext, &types.Membership{
		Project: result.Project,
		User:    result.User,
		Role:    enum.RoleDeveloper,
	})
	if err != nil {
		t.Error(err)
		return
	}

	updated, err := store.Find(noContext, result.Project, result.User)
	if err != nil {
		t.Error(err)
		return
	}

	if got, want := updated.Role, result.Role; got != want {
		t.Errorf("Want role %v, got %v", want, got)
	}
}

func TestMembershipDelete(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err := seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewMemberStoreSync(NewMemberStore(db))
	_, err = store.Find(noContext, 1, 1)
	if err != nil {
		t.Error(err)
	}

	err = store.Delete(noContext, 1, 1)
	if err != nil {
		t.Error(err)
	}

	_, err = store.Find(noContext, 1, 1)
	if err != sql.ErrNoRows {
		t.Errorf("Expected ErrNoRows, got %v", err)
	}
}

func TestMembershipUniqueIndex(t *testing.T) {
	t.Skip()
}

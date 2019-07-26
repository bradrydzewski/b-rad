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
)

func TestUserCount(t *testing.T) {
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

	users := NewUserStoreSync(NewUserStore(db))
	count, err := users.Count(noContext)
	if err != nil {
		t.Error(err)
	}
	if got, want := count, int64(2); got != want {
		t.Errorf("Want count %d, got %d", want, got)
	}
}

func TestUserFindID(t *testing.T) {
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

	users := NewUserStoreSync(NewUserStore(db))
	user, err := users.Find(noContext, 1)
	if err != nil {
		t.Error(err)
	}
	if got, want := user.Email, "jane@example.com"; want != got {
		t.Errorf("Want email %q, got %q", want, got)
	}
}

func TestUserFindEmail(t *testing.T) {
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

	users := NewUserStoreSync(NewUserStore(db))
	user, err := users.FindEmail(noContext, "jane@example.com")
	if err != nil {
		t.Error(err)
	}
	if got, want := user.Email, "jane@example.com"; want != got {
		t.Errorf("Want email %q, got %q", want, got)
	}
}

func TestUserFindToken(t *testing.T) {
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

	users := NewUserStoreSync(NewUserStore(db))
	user, err := users.FindToken(noContext, "12345")
	if err != nil {
		t.Error(err)
	}
	if got, want := user.Email, "jane@example.com"; want != got {
		t.Errorf("Want email %q, got %q", want, got)
	}
}

func TestUserFindEmailNocase(t *testing.T) {
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

	users := NewUserStoreSync(NewUserStore(db))
	user, err := users.FindEmail(noContext, "JANE@EXAMPLE.COM")
	if err != nil {
		t.Error(err)
	}
	if got, want := user.Email, "jane@example.com"; want != got {
		t.Errorf("Want email %q, got %q", want, got)
	}
}

func TestUserList(t *testing.T) {
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

	users := NewUserStoreSync(NewUserStore(db))
	results, err := users.List(noContext, types.Params{})
	if err != nil {
		t.Error(err)
	}
	if got, want := len(results), 2; got != want {
		t.Errorf("Want %d users, got %d", want, got)
	}
	if got, want := results[0].Email, "jane@example.com"; want != got {
		t.Errorf("Want email %q, got %q", want, got)
	}
	if got, want := results[1].Email, "john@example.com"; want != got {
		t.Errorf("Want email %q, got %q", want, got)
	}
}

func TestUserCreate(t *testing.T) {
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

	users := NewUserStoreSync(NewUserStore(db))
	user := &types.User{
		Email:   "jess@example.com",
		Token:   "8277e0910d750195b448797616e091ad",
		Admin:   true,
		Blocked: false,
		Created: 915148700,
		Updated: 915148800,
		Authed:  915148900,
	}
	if err := users.Create(noContext, user); err != nil {
		t.Error(err)
		return
	}
}

func TestUserUniqueIndexEmail(t *testing.T) {
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

	users := NewUserStoreSync(NewUserStore(db))
	user := &types.User{
		Email:   "jane@example.com",
		Token:   "8277e0910d750195b448797616e091ad",
		Admin:   true,
		Blocked: false,
		Created: 915148700,
		Updated: 915148800,
		Authed:  915148900,
	}
	if err := users.Create(noContext, user); err == nil {
		t.Errorf("Expect unique index violation")
	}
}

func TestUserUpdate(t *testing.T) {
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

	users := NewUserStoreSync(NewUserStore(db))
	user, err := users.Find(noContext, 1)
	if err != nil {
		t.Error(err)
	}

	user.Email = "noreply@example.com"
	err = users.Update(noContext, user)
	if err != nil {
		t.Error(err)
	}

	updated, err := users.Find(noContext, user.ID)
	if err != nil {
		t.Error(err)
	}

	if got, want := updated.Email, user.Email; got != want {
		t.Errorf("Want email %q, got %q", want, got)
	}
}

func TestUserDelete(t *testing.T) {
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

	users := NewUserStoreSync(NewUserStore(db))
	user, err := users.Find(noContext, 1)
	if err != nil {
		t.Error(err)
	}

	err = users.Delete(noContext, user)
	if err != nil {
		t.Error(err)
	}

	_, err = users.Find(noContext, 1)
	if err != sql.ErrNoRows {
		t.Errorf("Expected ErrNoRows, got %v", err)
	}
}

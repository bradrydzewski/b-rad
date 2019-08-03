// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/{{toLower repo}}/types/enum"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jmoiron/sqlx"
)

// member fields to ignore in test comparisons
var memberIgnore = cmpopts.IgnoreFields(types.Member{})

func TestMember(t *testing.T) {
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

	if _, err := newUserStoreSeeded(db); err != nil {
		t.Error(err)
		return
	}
	if _, err := new{{title project}}StoreSeeded(db); err != nil {
		t.Error(err)
		return
	}

	store := NewMemberStoreSync(NewMemberStore(db))
	t.Run("create", testMemberCreate(store))
	t.Run("find", testMemberFind(store))
	t.Run("list", testMemberList(store))
	t.Run("update", testMemberUpdate(store))
	t.Run("delete", testMemberDelete(store))
}

// this test creates entries in the database and confirms
// the primary keys were auto-incremented.
func testMemberCreate(store store.MemberStore) func(t *testing.T) {
	return func(t *testing.T) {
		vv := []*types.Membership{}
		if err := unmarshal("testdata/memberships.json", &vv); err != nil {
			t.Error(err)
			return
		}
		// create row 1
		v := vv[0]
		if err := store.Create(noContext, v); err != nil {
			t.Error(err)
			return
		}
		// create row 2
		v = vv[1]
		if err := store.Create(noContext, v); err != nil {
			t.Error(err)
			return
		}
		// create row 3
		v = vv[2]
		if err := store.Create(noContext, v); err != nil {
			t.Error(err)
			return
		}

		t.Run("duplicate", func(t *testing.T) {
			if err := store.Create(noContext, v); err == nil {
				t.Errorf("Expect duplicate row error")
				return
			}
		})
	}
}

// this test fetches members from the database by id
// and compares to the expected results (sourced from a json file)
// to ensure all columns are correctly mapped.
func testMemberFind(store store.MemberStore) func(t *testing.T) {
	return func(t *testing.T) {
		vv := []*types.Member{}
		if err := unmarshal("testdata/members.json", &vv); err != nil {
			t.Error(err)
			return
		}

		// Find row 1
		want := vv[0]
		got, err := store.Find(noContext, 1, 1)
		if err != nil {
			t.Error(err)
			return
		}
		if diff := cmp.Diff(got, want, memberIgnore); len(diff) != 0 {
			t.Errorf(diff)
			return
		}

		// Find row 2
		want = vv[1]
		got, err = store.Find(noContext, 1, 2)
		if err != nil {
			t.Error(err)
			return
		}
		if diff := cmp.Diff(got, want, memberIgnore); len(diff) != 0 {
			t.Errorf(diff)
			return
		}

		// Find row 3
		want = vv[2]
		got, err = store.Find(noContext, 2, 2)
		if err != nil {
			t.Error(err)
			return
		}
		if diff := cmp.Diff(got, want, memberIgnore); len(diff) != 0 {
			t.Errorf(diff)
			return
		}
	}
}

// this test fetches a list of members from the database
// and compares to the expected results (sourced from a json file)
// to ensure all columns are correctly mapped.
func testMemberList(store store.MemberStore) func(t *testing.T) {
	return func(t *testing.T) {
		want := []*types.Member{}
		if err := unmarshal("testdata/members.json", &want); err != nil {
			t.Error(err)
			return
		}
		got, err := store.List(noContext, 1, types.Params{Page: 0, Size: 100})
		if err != nil {
			t.Error(err)
			return
		}
		// take the first two items in the list
		// and ignore the third.
		want = want[:2]
		if diff := cmp.Diff(got, want, memberIgnore); len(diff) != 0 {
			t.Errorf(diff)
			return
		}
	}
}

// this test updates an member in the database and then fetches
// the member and confirms the column was updated as expected.
func testMemberUpdate(store store.MemberStore) func(t *testing.T) {
	return func(t *testing.T) {
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
			{{title project}}: result.{{title project}},
			User:    result.User,
			Role:    enum.RoleDeveloper,
		})
		if err != nil {
			t.Error(err)
			return
		}

		updated, err := store.Find(noContext, result.{{title project}}, result.User)
		if err != nil {
			t.Error(err)
			return
		}

		if got, want := updated.Role, result.Role; got != want {
			t.Errorf("Want role %v, got %v", want, got)
		}
	}
}

// this test deletes an member from the database and then confirms
// subsequent attempts to fetch the deleted member result in
// a sql.ErrNoRows error.
func testMemberDelete(store store.MemberStore) func(t *testing.T) {
	return func(t *testing.T) {
		v, err := store.Find(noContext, 1, 1)
		if err != nil {
			t.Error(err)
			return
		}
		if err := store.Delete(noContext, v.{{title project}}, v.User); err != nil {
			t.Error(err)
			return
		}
		if _, err := store.Find(noContext, v.{{title project}}, v.User); err != sql.ErrNoRows {
			t.Errorf("Expected sql.ErrNoRows got %s", err)
		}
	}
}

// helper function that returns a member store that is
// seeded with {{toLower project}} member data loaded from a json file.
func newMemberStoreSeeded(db *sqlx.DB) (store.MemberStore, error) {
	store := NewMemberStoreSync(NewMemberStore(db))
	vv := []*types.Membership{}
	if err := unmarshal("testdata/memberships.json", &vv); err != nil {
		return nil, err
	}
	for _, v := range vv {
		if err := store.Create(noContext, v); err != nil {
			return nil, fmt.Errorf("{{toLower project}}: %v, user: %v, error: %w", v.{{title project}}, v.User, err)
		}
	}
	return store, nil
}

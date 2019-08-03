// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jmoiron/sqlx"
)

// {{toLower project}} fields to ignore in test comparisons
var {{toLower project}}Ignore = cmpopts.IgnoreFields(types.{{title project}}{},
	"ID", "Token", "Created", "Updated")

func Test{{title project}}(t *testing.T) {
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

	store := New{{title project}}StoreSync(New{{title project}}Store(db))
	t.Run("create", test{{title project}}Create(store))
	t.Run("find", test{{title project}}Find(store))

	// seed the database with membership data for
	// the subsequent list unit tests.
	if _, err := newMemberStoreSeeded(db); err != nil {
		t.Error(err)
		return
	}

	t.Run("list", test{{title project}}List(store))
	t.Run("update", test{{title project}}Update(store))
	t.Run("delete", test{{title project}}Delete(store))
}

// this test creates entries in the database and confirms
// the primary keys were auto-incremented.
func test{{title project}}Create(store store.{{title project}}Store) func(t *testing.T) {
	return func(t *testing.T) {
		vv := []*types.{{title project}}{}
		if err := unmarshal("testdata/{{toLower project}}s.json", &vv); err != nil {
			t.Error(err)
			return
		}
		// create row 1
		v := vv[0]
		// generate a deterministic token for each
		// entry based on the hash of the email.
		v.Token = fmt.Sprintf("%x", v.Slug)
		if err := store.Create(noContext, v); err != nil {
			t.Error(err)
			return
		}
		if v.ID == 0 {
			t.Errorf("Want autoincremented primary key")
		}
		// create row 2
		v = vv[1]
		v.Token = fmt.Sprintf("%x", v.Slug)
		if err := store.Create(noContext, v); err != nil {
			t.Error(err)
			return
		}
		if v.ID == 0 {
			t.Errorf("Want autoincremented primary key")
		}

		t.Run("duplicate slug", func(t *testing.T) {
			v.ID = 0
			v.Token = "9afeab83324a53"
			v.Slug = "cassini"
			if err := store.Create(noContext, v); err == nil {
				t.Errorf("Expect duplicate row error")
				return
			}
		})

		t.Run("duplicate token", func(t *testing.T) {
			v.ID = 0
			v.Slug = "voyager"
			v.Token = "63617373696e69"
			if err := store.Create(noContext, v); err == nil {
				t.Errorf("Expect duplicate row error")
				return
			}
		})
	}
}

// this test fetches {{toLower project}}s from the database by id and key
// and compares to the expected results (sourced from a json file)
// to ensure all columns are correctly mapped.
func test{{title project}}Find(store store.{{title project}}Store) func(t *testing.T) {
	return func(t *testing.T) {
		vv := []*types.{{title project}}{}
		if err := unmarshal("testdata/{{toLower project}}s.json", &vv); err != nil {
			t.Error(err)
			return
		}
		want := vv[0]
		want.Token = "63617373696e69"

		// Find row by ID
		got, err := store.Find(noContext, 1)
		if err != nil {
			t.Error(err)
			return
		}
		if diff := cmp.Diff(got, want, {{toLower project}}Ignore); len(diff) != 0 {
			t.Errorf(diff)
			return
		}

		t.Run("token", func(t *testing.T) {
			got, err := store.FindToken(noContext, want.Token)
			if err != nil {
				t.Error(err)
				return
			}
			if diff := cmp.Diff(got, want, {{toLower project}}Ignore); len(diff) != 0 {
				t.Errorf(diff)
				return
			}
		})

		t.Run("slug", func(t *testing.T) {
			got, err := store.FindSlug(noContext, want.Slug)
			if err != nil {
				t.Error(err)
				return
			}
			if diff := cmp.Diff(got, want, {{toLower project}}Ignore); len(diff) != 0 {
				t.Errorf(diff)
				return
			}
		})

		t.Run("slug", func(t *testing.T) {
			got, err := store.FindSlug(noContext, want.Slug)
			if err != nil {
				t.Error(err)
				return
			}
			if diff := cmp.Diff(got, want, {{toLower project}}Ignore); len(diff) != 0 {
				t.Errorf(diff)
				return
			}
		})
	}
}

// this test fetches a list of {{toLower project}}s from the database
// and compares to the expected results (sourced from a json file)
// to ensure all columns are correctly mapped.
func test{{title project}}List(store store.{{title project}}Store) func(t *testing.T) {
	return func(t *testing.T) {
		want := []*types.{{title project}}{}
		if err := unmarshal("testdata/{{toLower project}}s.json", &want); err != nil {
			t.Error(err)
			return
		}
		got, err := store.List(noContext, 2, types.Params{Page: 0, Size: 100})
		if err != nil {
			t.Error(err)
			return
		}
		if len(got) != 2 {
			t.Errorf("Expect 2 {{toLower project}}s")
		}
		if diff := cmp.Diff(got, want, {{toLower project}}Ignore); len(diff) != 0 {
			t.Errorf(diff)
			return
		}
	}
}

// this test updates an {{toLower project}} in the database and then fetches
// the {{toLower project}} and confirms the column was updated as expected.
func test{{title project}}Update(store store.{{title project}}Store) func(t *testing.T) {
	return func(t *testing.T) {
		before, err := store.Find(noContext, 1)
		if err != nil {
			t.Error(err)
			return
		}
		before.Updated = time.Now().Unix()
		before.Active = false
		if err := store.Update(noContext, before); err != nil {
			t.Error(err)
			return
		}
		after, err := store.Find(noContext, 1)
		if err != nil {
			t.Error(err)
			return
		}

		if diff := cmp.Diff(before, after, {{toLower project}}Ignore); len(diff) != 0 {
			t.Errorf(diff)
			return
		}
	}
}

// this test deletes an {{toLower project}} from the database and then confirms
// subsequent attempts to fetch the deleted {{toLower project}} result in
// a sql.ErrNoRows error.
func test{{title project}}Delete(store store.{{title project}}Store) func(t *testing.T) {
	return func(t *testing.T) {
		v, err := store.Find(noContext, 1)
		if err != nil {
			t.Error(err)
			return
		}
		if err := store.Delete(noContext, v); err != nil {
			t.Error(err)
			return
		}
		if _, err := store.Find(noContext, 1); err != sql.ErrNoRows {
			t.Errorf("Expected sql.ErrNoRows got %s", err)
		}
	}
}

// helper function that returns an {{toLower project}} store that is seeded
// with {{toLower project}} data loaded from a json file.
func new{{title project}}StoreSeeded(db *sqlx.DB) (store.{{title project}}Store, error) {
	store := New{{title project}}StoreSync(New{{title project}}Store(db))
	vv := []*types.{{title project}}{}
	if err := unmarshal("testdata/{{toLower project}}s.json", &vv); err != nil {
		return nil, err
	}
	for _, v := range vv {
		v.Token = fmt.Sprintf("%x", v.Slug)
		if err := store.Create(noContext, v); err != nil {
			return nil, err
		}
	}
	return store, nil
}

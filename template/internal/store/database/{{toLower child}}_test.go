// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"database/sql"
	"testing"

	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jmoiron/sqlx"
)

// {{toLower child}} fields to ignore in test comparisons
var {{toLower child}}Ignore = cmpopts.IgnoreFields(types.{{title child}}{},
	"ID", "Created", "Updated")

func Test{{title child}}(t *testing.T) {
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

	if _, err := new{{title project}}StoreSeeded(db); err != nil {
		t.Error(err)
		return
	}
	if _, err := new{{title parent}}StoreSeeded(db); err != nil {
		t.Error(err)
		return
	}

	store := New{{title child}}StoreSync(New{{title child}}Store(db))
	t.Run("create", test{{title child}}Create(store))
	t.Run("find", test{{title child}}Find(store))
	t.Run("list", test{{title child}}List(store))
	t.Run("update", test{{title child}}Update(store))
	t.Run("delete", test{{title child}}Delete(store))
}

// this test creates entries in the database and confirms
// the primary keys were auto-incremented.
func test{{title child}}Create(store store.{{title child}}Store) func(t *testing.T) {
	return func(t *testing.T) {
		vv := []*types.{{title child}}{}
		if err := unmarshal("testdata/{{toLower child}}s.json", &vv); err != nil {
			t.Error(err)
			return
		}

		// create row 1
		v := vv[0]
		if err := store.Create(noContext, v); err != nil {
			t.Error(err)
			return
		}
		if v.ID == 0 {
			t.Errorf("Want autoincremented primary key")
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

		t.Run("duplicate slug", func(t *testing.T) {
			// reset the ID so that a new row is created
			// using the same slug
			v.ID = 0
			if err := store.Create(noContext, v); err == nil {
				t.Errorf("Expect duplicate row error")
				return
			}
		})
	}
}

// this test fetches {{toLower child}}s from the database by id and key
// and compares to the expected results (sourced from a json file)
// to ensure all columns are correctly mapped.
func test{{title child}}Find(store store.{{title child}}Store) func(t *testing.T) {
	return func(t *testing.T) {
		vv := []*types.{{title child}}{}
		if err := unmarshal("testdata/{{toLower child}}s.json", &vv); err != nil {
			t.Error(err)
			return
		}
		want := vv[0]

		t.Run("id", func(t *testing.T) {
			got, err := store.Find(noContext, 1)
			if err != nil {
				t.Error(err)
				return
			}
			if diff := cmp.Diff(got, want, {{toLower child}}Ignore); len(diff) != 0 {
				t.Errorf(diff)
				return
			}
		})

		t.Run("slug", func(t *testing.T) {
			got, err := store.FindSlug(noContext, want.{{title parent}}, want.Slug)
			if err != nil {
				t.Error(err)
				return
			}
			if diff := cmp.Diff(got, want, {{toLower child}}Ignore); len(diff) != 0 {
				t.Errorf(diff)
				return
			}
		})
	}
}

// this test fetches a list of {{toLower child}}s from the database
// and compares to the expected results (sourced from a json file)
// to ensure all columns are correctly mapped.
func test{{title child}}List(store store.{{title child}}Store) func(t *testing.T) {
	return func(t *testing.T) {
		want := []*types.{{title child}}{}
		if err := unmarshal("testdata/{{toLower child}}s.json", &want); err != nil {
			t.Error(err)
			return
		}
		got, err := store.List(noContext, 2, types.Params{Size: 25, Page: 0})
		if err != nil {
			t.Error(err)
			return
		}

		if diff := cmp.Diff(got, want[1:], {{toLower child}}Ignore); len(diff) != 0 {
			t.Errorf(diff)
			debug(t, got)
			return
		}
	}
}

// this test updates a {{toLower child}} in the database and then fetches
// the {{toLower child}} and confirms the column was updated as expected.
func test{{title child}}Update(store store.{{title child}}Store) func(t *testing.T) {
	return func(t *testing.T) {
		before, err := store.Find(noContext, 1)
		if err != nil {
			t.Error(err)
			return
		}
		before.Desc = "updated description"
		if err := store.Update(noContext, before); err != nil {
			t.Error(err)
			return
		}
		after, err := store.Find(noContext, 1)
		if err != nil {
			t.Error(err)
			return
		}

		if diff := cmp.Diff(before, after, {{toLower child}}Ignore); len(diff) != 0 {
			t.Errorf(diff)
			return
		}
	}
}

// this test deletes a {{toLower child}} from the database and then confirms
// subsequent attempts to fetch the deleted {{toLower child}} result in
// a sql.ErrNoRows error.
func test{{title child}}Delete(store store.{{title child}}Store) func(t *testing.T) {
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

// helper function that returns an {{toLower child}} store that is seeded
// with {{toLower child}} data loaded from a json file.
func new{{title child}}StoreSeeded(db *sqlx.DB) (store.{{title child}}Store, error) {
	store := New{{title child}}StoreSync(New{{title child}}Store(db))
	vv := []*types.{{title child}}{}
	if err := unmarshal("testdata/{{toLower child}}s.json", &vv); err != nil {
		return nil, err
	}
	for _, v := range vv {
		if err := store.Create(noContext, v); err != nil {
			return nil, err
		}
	}
	return store, nil
}

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

// {{toLower parent}} fields to ignore in test comparisons
var {{toLower parent}}Ignore = cmpopts.IgnoreFields(types.{{title parent}}{},
	"ID", "Created", "Updated")

func Test{{title parent}}(t *testing.T) {
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

	store := New{{title parent}}StoreSync(New{{title parent}}Store(db))
	t.Run("create", test{{title parent}}Create(store))
	t.Run("find", test{{title parent}}Find(store))
	t.Run("list", test{{title parent}}List(store))
	t.Run("update", test{{title parent}}Update(store))
	t.Run("delete", test{{title parent}}Delete(store))
}

// this test creates entries in the database and confirms
// the primary keys were auto-incremented.
func test{{title parent}}Create(store store.{{title parent}}Store) func(t *testing.T) {
	return func(t *testing.T) {
		vv := []*types.{{title parent}}{}
		if err := unmarshal("testdata/{{toLower parent}}s.json", &vv); err != nil {
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

// this test fetches {{toLower parent}}s from the database by id and key
// and compares to the expected results (sourced from a json file)
// to ensure all columns are correctly mapped.
func test{{title parent}}Find(store store.{{title parent}}Store) func(t *testing.T) {
	return func(t *testing.T) {
		vv := []*types.{{title parent}}{}
		if err := unmarshal("testdata/{{toLower parent}}s.json", &vv); err != nil {
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
			if diff := cmp.Diff(got, want, {{toLower parent}}Ignore); len(diff) != 0 {
				t.Errorf(diff)
				return
			}
		})

		t.Run("slug", func(t *testing.T) {
			got, err := store.FindSlug(noContext, want.{{title project}}, want.Slug)
			if err != nil {
				t.Error(err)
				return
			}
			if diff := cmp.Diff(got, want, {{toLower parent}}Ignore); len(diff) != 0 {
				t.Errorf(diff)
				return
			}
		})
	}
}

// this test fetches a list of {{toLower parent}}s from the database
// and compares to the expected results (sourced from a json file)
// to ensure all columns are correctly mapped.
func test{{title parent}}List(store store.{{title parent}}Store) func(t *testing.T) {
	return func(t *testing.T) {
		want := []*types.{{title parent}}{}
		if err := unmarshal("testdata/{{toLower parent}}s.json", &want); err != nil {
			t.Error(err)
			return
		}
		got, err := store.List(noContext, 2, types.Params{Size: 25, Page: 0})
		if err != nil {
			t.Error(err)
			return
		}

		if diff := cmp.Diff(got, want[1:], {{toLower parent}}Ignore); len(diff) != 0 {
			t.Errorf(diff)
			debug(t, got)
			return
		}
	}
}

// this test updates a {{toLower parent}} in the database and then fetches
// the {{toLower parent}} and confirms the column was updated as expected.
func test{{title parent}}Update(store store.{{title parent}}Store) func(t *testing.T) {
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

		if diff := cmp.Diff(before, after, {{toLower parent}}Ignore); len(diff) != 0 {
			t.Errorf(diff)
			return
		}
	}
}

// this test deletes a {{toLower parent}} from the database and then confirms
// subsequent attempts to fetch the deleted {{toLower parent}} result in
// a sql.ErrNoRows error.
func test{{title parent}}Delete(store store.{{title parent}}Store) func(t *testing.T) {
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

// helper function that returns an {{toLower parent}} store that is seeded
// with {{toLower parent}} data loaded from a json file.
func new{{title parent}}StoreSeeded(db *sqlx.DB) (store.{{title parent}}Store, error) {
	store := New{{title parent}}StoreSync(New{{title parent}}Store(db))
	vv := []*types.{{title parent}}{}
	if err := unmarshal("testdata/{{toLower parent}}s.json", &vv); err != nil {
		return nil, err
	}
	for _, v := range vv {
		if err := store.Create(noContext, v); err != nil {
			return nil, err
		}
	}
	return store, nil
}

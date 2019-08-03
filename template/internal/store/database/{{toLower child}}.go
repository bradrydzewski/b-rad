// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"context"

	"github.com/{{toLower repo}}/internal/store"
	"github.com/{{toLower repo}}/types"

	"github.com/jmoiron/sqlx"
)

var _ store.{{title child}}Store = (*{{title child}}Store)(nil)

// New{{title child}}Store returns a new {{title child}}Store.
func New{{title child}}Store(db *sqlx.DB) *{{title child}}Store {
	return &{{title child}}Store{db}
}

// {{title child}}Store implements a {{title child}}Store backed by a relational
// database.
type {{title child}}Store struct {
	db *sqlx.DB
}

// Find finds the {{toLower child}} by id.
func (s *{{title child}}Store) Find(ctx context.Context, id int64) (*types.{{title child}}, error) {
	dst := new(types.{{title child}})
	err := s.db.Get(dst, {{toLower child}}SelectID, id)
	return dst, err
}

// FindSlug finds the {{toLower child}} by {{toLower parent}} id and slug.
func (s *{{title child}}Store) FindSlug(ctx context.Context, id int64, slug string) (*types.{{title child}}, error) {
	dst := new(types.{{title child}})
	err := s.db.Get(dst, {{toLower child}}SelectSlug, id, slug)
	return dst, err
}

// List returns a list of {{toLower child}}s.
func (s *{{title child}}Store) List(ctx context.Context, id int64, opts types.Params) ([]*types.{{title child}}, error) {
	dst := []*types.{{title child}}{}
	err := s.db.Select(&dst, {{toLower child}}Select, id, limit(opts.Size), offset(opts.Page, opts.Size))
	return dst, err
}

// Create saves the {{toLower child}} details.
func (s *{{title child}}Store) Create(ctx context.Context, {{toLower child}} *types.{{title child}}) error {
	query, arg, err := s.db.BindNamed({{toLower child}}Insert, {{toLower child}})
	if err != nil {
		return err
	}
	return s.db.QueryRow(query, arg...).Scan(&{{toLower child}}.ID)
}

// Update updates the {{toLower child}} details.
func (s *{{title child}}Store) Update(ctx context.Context, {{toLower child}} *types.{{title child}}) error {
	query, arg, err := s.db.BindNamed({{toLower child}}Update, {{toLower child}})
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, arg...)
	return err
}

// Delete deletes the {{toLower child}}.
func (s *{{title child}}Store) Delete(ctx context.Context, {{toLower child}} *types.{{title child}}) error {
	_, err := s.db.Exec({{toLower child}}Delete, {{toLower child}}.ID)
	return err
}

const {{toLower child}}Base = `
SELECT
 {{toLower child}}_id
,{{toLower child}}_{{toLower project}}_id
,{{toLower child}}_{{toLower parent}}_id
,{{toLower child}}_slug
,{{toLower child}}_name
,{{toLower child}}_desc
,{{toLower child}}_created
,{{toLower child}}_updated
FROM {{toLower child}}s
`

const {{toLower child}}Select = {{toLower child}}Base + `
WHERE {{toLower child}}_{{toLower parent}}_id = $1
ORDER BY {{toLower child}}_name ASC
LIMIT $2 OFFSET $3
`

const {{toLower child}}SelectID = {{toLower child}}Base + `
WHERE {{toLower child}}_id = $1
`

const {{toLower child}}SelectSlug = {{toLower child}}Base + `
WHERE {{toLower child}}_{{toLower parent}}_id = $1
  AND {{toLower child}}_slug     = $2
`

const {{toLower child}}Insert = `
INSERT INTO {{toLower child}}s (
 {{toLower child}}_{{toLower project}}_id
,{{toLower child}}_{{toLower parent}}_id
,{{toLower child}}_slug
,{{toLower child}}_name
,{{toLower child}}_desc
,{{toLower child}}_created
,{{toLower child}}_updated
) values (
 :{{toLower child}}_{{toLower project}}_id
,:{{toLower child}}_{{toLower parent}}_id
,:{{toLower child}}_slug
,:{{toLower child}}_name
,:{{toLower child}}_desc
,:{{toLower child}}_created
,:{{toLower child}}_updated
) RETURNING {{toLower child}}_id
`

const {{toLower child}}Update = `
UPDATE {{toLower child}}s
SET
 {{toLower child}}_name    = :{{toLower child}}_name
,{{toLower child}}_desc    = :{{toLower child}}_desc
,{{toLower child}}_updated = :{{toLower child}}_updated
WHERE {{toLower child}}_id = :{{toLower child}}_id
`

const {{toLower child}}Delete = `
DELETE FROM {{toLower child}}s
WHERE {{toLower child}}_id = $1
`

const {{toLower child}}Delete{{title parent}} = `
DELETE FROM {{toLower child}}s
WHERE {{toLower child}}_{{toLower parent}}_id = $1
`

const {{toLower child}}Delete{{title project}} = `
DELETE FROM {{toLower child}}s
WHERE {{toLower child}}_{{toLower project}}_id = $1
`

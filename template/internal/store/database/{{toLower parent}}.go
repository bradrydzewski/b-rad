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

var _ store.{{title parent}}Store = (*{{title parent}}Store)(nil)

// New{{title parent}}Store returns a new {{title parent}}Store.
func New{{title parent}}Store(db *sqlx.DB) *{{title parent}}Store {
	return &{{title parent}}Store{db}
}

// {{title parent}}Store implements a {{title parent}}Store backed by a relational
// database.
type {{title parent}}Store struct {
	db *sqlx.DB
}

// Find finds the {{toLower parent}} by id.
func (s *{{title parent}}Store) Find(ctx context.Context, id int64) (*types.{{title parent}}, error) {
	dst := new(types.{{title parent}})
	err := s.db.Get(dst, {{toLower parent}}SelectID, id)
	return dst, err
}

// FindSlug finds the {{toLower parent}} by {{toLower project}} id and slug.
func (s *{{title parent}}Store) FindSlug(ctx context.Context, id int64, slug string) (*types.{{title parent}}, error) {
	dst := new(types.{{title parent}})
	err := s.db.Get(dst, {{toLower parent}}SelectSlug, id, slug)
	return dst, err
}

// List returns a list of {{toLower parent}}s.
func (s *{{title parent}}Store) List(ctx context.Context, id int64, opts types.Params) ([]*types.{{title parent}}, error) {
	dst := []*types.{{title parent}}{}
	err := s.db.Select(&dst, {{toLower parent}}Select, id, limit(opts.Size), offset(opts.Page, opts.Size))
	return dst, err
}

// Create saves the {{toLower parent}} details.
func (s *{{title parent}}Store) Create(ctx context.Context, {{toLower parent}} *types.{{title parent}}) error {
	query, arg, err := s.db.BindNamed({{toLower parent}}Insert, {{toLower parent}})
	if err != nil {
		return err
	}
	return s.db.QueryRow(query, arg...).Scan(&{{toLower parent}}.ID)
}

// Update updates the {{toLower parent}} details.
func (s *{{title parent}}Store) Update(ctx context.Context, {{toLower parent}} *types.{{title parent}}) error {
	query, arg, err := s.db.BindNamed({{toLower parent}}Update, {{toLower parent}})
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, arg...)
	return err
}

// Delete deletes the {{toLower parent}}.
func (s *{{title parent}}Store) Delete(ctx context.Context, {{toLower parent}} *types.{{title parent}}) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	// delete the {{toLower child}}s associated with the {{toLower parent}}
	if _, err := tx.Exec({{toLower child}}Delete{{title parent}}, {{toLower parent}}.ID); err != nil {
		return err
	}
	// delete the {{toLower parent}}
	if _, err := tx.Exec({{toLower parent}}Delete, {{toLower parent}}.ID); err != nil {
		return err
	}
	return tx.Commit()
}

const {{toLower parent}}Base = `
SELECT
 {{toLower parent}}_id
,{{toLower parent}}_{{toLower project}}_id
,{{toLower parent}}_slug
,{{toLower parent}}_name
,{{toLower parent}}_desc
,{{toLower parent}}_created
,{{toLower parent}}_updated
FROM {{toLower parent}}s
`

const {{toLower parent}}Select = {{toLower parent}}Base + `
WHERE {{toLower parent}}_{{toLower project}}_id = $1
ORDER BY {{toLower parent}}_name ASC
LIMIT $2 OFFSET $3
`

const {{toLower parent}}SelectID = {{toLower parent}}Base + `
WHERE {{toLower parent}}_id = $1
`

const {{toLower parent}}SelectSlug = {{toLower parent}}Base + `
WHERE {{toLower parent}}_{{toLower project}}_id = $1
  AND {{toLower parent}}_slug = $2
`

const {{toLower parent}}Insert = `
INSERT INTO {{toLower parent}}s (
 {{toLower parent}}_{{toLower project}}_id
,{{toLower parent}}_slug
,{{toLower parent}}_name
,{{toLower parent}}_desc
,{{toLower parent}}_created
,{{toLower parent}}_updated
) values (
 :{{toLower parent}}_{{toLower project}}_id
,:{{toLower parent}}_slug
,:{{toLower parent}}_name
,:{{toLower parent}}_desc
,:{{toLower parent}}_created
,:{{toLower parent}}_updated
) RETURNING {{toLower parent}}_id
`

const {{toLower parent}}Update = `
UPDATE {{toLower parent}}s
SET
 {{toLower parent}}_name    = :{{toLower parent}}_name
,{{toLower parent}}_desc    = :{{toLower parent}}_desc
,{{toLower parent}}_updated = :{{toLower parent}}_updated
WHERE {{toLower parent}}_id = :{{toLower parent}}_id
`

const {{toLower parent}}Delete = `
DELETE FROM {{toLower parent}}s
WHERE {{toLower parent}}_id = $1
`

const {{toLower parent}}Delete{{title project}} = `
DELETE FROM {{toLower parent}}s
WHERE {{toLower parent}}_{{toLower project}}_id = $1
`

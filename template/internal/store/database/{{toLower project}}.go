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

var _ store.{{title project}}Store = (*{{title project}}Store)(nil)

// New{{title project}}Store returns a new {{title project}}tStore.
func New{{title project}}Store(db *sqlx.DB) *{{title project}}Store {
	return &{{title project}}Store{db}
}

// {{title project}}Store implements a {{title project}}Store backed by a
// relational database.
type {{title project}}Store struct {
	db *sqlx.DB
}

// Find finds the {{toLower project}} by id.
func (s *{{title project}}Store) Find(ctx context.Context, id int64) (*types.{{title project}}, error) {
	dst := new(types.{{title project}})
	err := s.db.Get(dst, {{toLower project}}SelectID, id)
	return dst, err
}

// FindToken finds the {{toLower project}} by token.
func (s *{{title project}}Store) FindToken(ctx context.Context, token string) (*types.{{title project}}, error) {
	dst := new(types.{{title project}})
	err := s.db.Get(dst, {{toLower project}}SelectToken, token)
	return dst, err
}

// List returns a list of {{toLower project}}s by user.
func (s *{{title project}}Store) List(ctx context.Context, user int64, opts types.Params) ([]*types.{{title project}}, error) {
	dst := []*types.{{title project}}{}
	err := s.db.Select(&dst, {{toLower project}}Select, user)
	// TODO(bradrydzewski) add limit and offset
	return dst, err
}

// Create saves the {{toLower project}} details.
func (s *{{title project}}Store) Create(ctx context.Context, {{toLower project}} *types.{{title project}}) error {
	query := {{toLower project}}Insert

	if s.db.DriverName() == "postgres" {
		query = {{toLower project}}InsertPg
	}

	query, arg, err := s.db.BindNamed(query, {{toLower project}})
	if err != nil {
		return err
	}

	if s.db.DriverName() == "postgres" {
		return s.db.QueryRow(query, arg...).Scan(&{{toLower project}}.ID)
	}

	res, err := s.db.Exec(query, arg...)
	if err != nil {
		return err
	}
	{{toLower project}}.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// Update updates the {{toLower project}} details.
func (s *{{title project}}Store) Update(ctx context.Context, {{toLower project}} *types.{{title project}}) error {
	query, arg, err := s.db.BindNamed({{toLower project}}Update, {{toLower project}})
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, arg...)
	return err
}

// Delete deletes the {{toLower project}}.
func (s *{{title project}}Store) Delete(ctx context.Context, {{toLower project}} *types.{{title project}}) error {
	_, err := s.db.Exec({{toLower project}}Delete, {{toLower project}}.ID)
	return err
}

const {{toLower project}}Base = `
SELECT
 {{toLower project}}_id
,{{toLower project}}_name
,{{toLower project}}_desc
,{{toLower project}}_token
,{{toLower project}}_active
,{{toLower project}}_created
,{{toLower project}}_updated
FROM {{toLower project}}s
`

const {{toLower project}}Select = {{toLower project}}Base + `
WHERE {{toLower project}}_id IN (
  SELECT member_{{toLower project}}_id
  FROM members
  WHERE member_user_id = $1
)
ORDER BY {{toLower project}}_name
`

const {{toLower project}}SelectID = {{toLower project}}Base + `
WHERE {{toLower project}}_id = $1
`

const {{toLower project}}SelectToken = {{toLower project}}Base + `
WHERE {{toLower project}}_token = $1
`

const {{toLower project}}Delete = `
DELETE FROM {{toLower project}}s
WHERE {{toLower project}}_id = $1
`

const {{toLower project}}Insert = `
INSERT INTO {{toLower project}}s (
 {{toLower project}}_name
,{{toLower project}}_desc
,{{toLower project}}_token
,{{toLower project}}_active
,{{toLower project}}_created
,{{toLower project}}_updated
) values (
 :{{toLower project}}_name
,:{{toLower project}}_desc
,:{{toLower project}}_token
,:{{toLower project}}_active
,:{{toLower project}}_created
,:{{toLower project}}_updated
)
`

const {{toLower project}}InsertPg = {{toLower project}}Insert + `
RETURNING {{toLower project}}_id
`

const {{toLower project}}Update = `
UPDATE {{toLower project}}s
SET
 {{toLower project}}_name      = :{{toLower project}}_name
,{{toLower project}}_desc      = :{{toLower project}}_desc
,{{toLower project}}_active    = :{{toLower project}}_active
,{{toLower project}}_updated   = :{{toLower project}}_updated
WHERE {{toLower project}}_id = :{{toLower project}}_id
`

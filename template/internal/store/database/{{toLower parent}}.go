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
	"github.com/{{github}}/types"

	"github.com/jmoiron/sqlx"
)

var _ store.{{parent}}Store = (*{{parent}}Store)(nil)

// New{{parent}}Store returns a new {{parent}}Store.
func New{{parent}}Store(db *sqlx.DB) *{{parent}}Store {
	return &{{parent}}Store{db}
}

// {{parent}}Store implements a {{parent}}Store backed by a relational
// database.
type {{parent}}Store struct {
	db *sqlx.DB
}

// Find finds the {{toLower parent}} by id.
func (s *{{parent}}Store) Find(ctx context.Context, id int64) (*types.{{parent}}, error) {
	dst := new(types.{{parent}})
	err := s.db.Get(dst, {{toLower parent}}SelectID, id)
	return dst, err
}

// List returns a list of {{toLower parent}}s.
func (s *{{parent}}Store) List(ctx context.Context, id int64, opts types.Params) ([]*types.{{parent}}, error) {
	dst := []*types.{{parent}}{}
	err := s.db.Select(&dst, {{toLower parent}}Select, id)
	// TODO(bradrydzewski) add limit and offset
	return dst, err
}

// Create saves the {{toLower parent}} details.
func (s *{{parent}}Store) Create(ctx context.Context, {{toLower parent}} *types.{{parent}}) error {
	query := {{toLower parent}}Insert

	if s.db.DriverName() == "postgres" {
		query = {{toLower parent}}InsertPg
	}

	query, arg, err := s.db.BindNamed(query, {{toLower parent}})
	if err != nil {
		return err
	}

	if s.db.DriverName() == "postgres" {
		return s.db.QueryRow(query, arg...).Scan(&{{toLower parent}}.ID)
	}

	res, err := s.db.Exec(query, arg...)
	if err != nil {
		return err
	}
	{{toLower parent}}.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// Update updates the {{toLower parent}} details.
func (s *{{parent}}Store) Update(ctx context.Context, {{toLower parent}} *types.{{parent}}) error {
	query, arg, err := s.db.BindNamed({{toLower parent}}Update, {{toLower parent}})
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, arg...)
	return err
}

// Delete deletes the {{toLower parent}}.
func (s *{{parent}}Store) Delete(ctx context.Context, {{toLower parent}} *types.{{parent}}) error {
	_, err := s.db.Exec({{toLower parent}}Delete, {{toLower parent}}.ID)
	return err
}

const {{toLower parent}}Base = `
SELECT
 {{toLower parent}}_id
,{{toLower parent}}_project_id
,{{toLower parent}}_name
,{{toLower parent}}_desc
,{{toLower parent}}_created
,{{toLower parent}}_updated
FROM {{toLower parent}}s
`

const {{toLower parent}}Select = {{toLower parent}}Base + `
WHERE {{toLower parent}}_project_id = $1
ORDER BY {{toLower parent}}_name ASC
`

const {{toLower parent}}SelectID = {{toLower parent}}Base + `
WHERE {{toLower parent}}_id = $1
`

const {{toLower parent}}Delete = `
DELETE FROM {{toLower parent}}s
WHERE {{toLower parent}}_id = $1
`

const {{toLower parent}}Insert = `
INSERT INTO {{toLower parent}}s (
 {{toLower parent}}_project_id
,{{toLower parent}}_name
,{{toLower parent}}_desc
,{{toLower parent}}_created
,{{toLower parent}}_updated
) values (
 :{{toLower parent}}_project_id
,:{{toLower parent}}_name
,:{{toLower parent}}_desc
,:{{toLower parent}}_created
,:{{toLower parent}}_updated
)
`

const {{toLower parent}}InsertPg = {{toLower parent}}Insert + `
RETURNING {{toLower parent}}_id
`

const {{toLower parent}}Update = `
UPDATE {{toLower parent}}s
SET
 {{toLower parent}}_name    = :{{toLower parent}}_name
,{{toLower parent}}_desc    = :{{toLower parent}}_desc
,{{toLower parent}}_updated = :{{toLower parent}}_updated
WHERE {{toLower parent}}_id = :{{toLower parent}}_id
`

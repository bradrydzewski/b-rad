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

var _ store.{{child}}Store = (*{{child}}Store)(nil)

// New{{child}}Store returns a new {{child}}Store.
func New{{child}}Store(db *sqlx.DB) *{{child}}Store {
	return &{{child}}Store{db}
}

// {{child}}Store implements a {{child}}Store backed by a relational
// database.
type {{child}}Store struct {
	db *sqlx.DB
}

// Find finds the {{toLower child}} by id.
func (s *{{child}}Store) Find(ctx context.Context, id int64) (*types.{{child}}, error) {
	dst := new(types.{{child}})
	err := s.db.Get(dst, {{toLower child}}SelectID, id)
	return dst, err
}

// List returns a list of {{toLower child}}s.
func (s *{{child}}Store) List(ctx context.Context, id int64, opts types.Params) ([]*types.{{child}}, error) {
	dst := []*types.{{child}}{}
	err := s.db.Select(&dst, {{toLower child}}Select, id)
	// TODO(bradrydzewski) add limit and offset
	return dst, err
}

// Create saves the {{toLower child}} details.
func (s *{{child}}Store) Create(ctx context.Context, {{toLower child}} *types.{{child}}) error {
	query := {{toLower child}}Insert

	if s.db.DriverName() == "postgres" {
		query = {{toLower child}}InsertPg
	}

	query, arg, err := s.db.BindNamed(query, {{toLower child}})
	if err != nil {
		return err
	}

	if s.db.DriverName() == "postgres" {
		return s.db.QueryRow(query, arg...).Scan(&{{toLower child}}.ID)
	}

	res, err := s.db.Exec(query, arg...)
	if err != nil {
		return err
	}
	{{toLower child}}.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// Update updates the {{toLower child}} details.
func (s *{{child}}Store) Update(ctx context.Context, {{toLower child}} *types.{{child}}) error {
	query, arg, err := s.db.BindNamed({{toLower child}}Update, {{toLower child}})
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, arg...)
	return err
}

// Delete deletes the {{toLower child}}.
func (s *{{child}}Store) Delete(ctx context.Context, {{toLower child}} *types.{{child}}) error {
	_, err := s.db.Exec({{toLower child}}Delete, {{toLower child}}.ID)
	return err
}

const {{toLower child}}Base = `
SELECT
 {{toLower child}}_id
,{{toLower child}}_{{toLower parent}}_id
,{{toLower child}}_name
,{{toLower child}}_desc
,{{toLower child}}_created
,{{toLower child}}_updated
FROM {{toLower child}}s
`

const {{toLower child}}Select = {{toLower child}}Base + `
WHERE {{toLower child}}_{{toLower parent}}_id = $1
ORDER BY {{toLower child}}_name ASC
`

const {{toLower child}}SelectID = {{toLower child}}Base + `
WHERE {{toLower child}}_id = $1
`

const {{toLower child}}Delete = `
DELETE FROM {{toLower child}}s
WHERE {{toLower child}}_id = $1
`

const {{toLower child}}Insert = `
INSERT INTO {{toLower child}}s (
 {{toLower child}}_{{toLower parent}}_id
,{{toLower child}}_name
,{{toLower child}}_desc
,{{toLower child}}_created
,{{toLower child}}_updated
) values (
 :{{toLower child}}_{{toLower parent}}_id
,:{{toLower child}}_name
,:{{toLower child}}_desc
,:{{toLower child}}_created
,:{{toLower child}}_updated
)
`

const {{toLower child}}InsertPg = {{toLower child}}Insert + `
RETURNING {{toLower child}}_id
`

const {{toLower child}}Update = `
UPDATE {{toLower child}}s
SET
 {{toLower child}}_name    = :{{toLower child}}_name
,{{toLower child}}_desc    = :{{toLower child}}_desc
,{{toLower child}}_updated = :{{toLower child}}_updated
WHERE {{toLower child}}_id = :{{toLower child}}_id
`

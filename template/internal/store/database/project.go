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

var _ store.ProjectStore = (*ProjectStore)(nil)

// NewProjectStore returns a new ProjecttStore.
func NewProjectStore(db *sqlx.DB) *ProjectStore {
	return &ProjectStore{db}
}

// ProjectStore implements a ProjectStore backed by a
// relational database.
type ProjectStore struct {
	db *sqlx.DB
}

// Find finds the project by id.
func (s *ProjectStore) Find(ctx context.Context, id int64) (*types.Project, error) {
	dst := new(types.Project)
	err := s.db.Get(dst, projectSelectID, id)
	return dst, err
}

// FindToken finds the project by token.
func (s *ProjectStore) FindToken(ctx context.Context, token string) (*types.Project, error) {
	dst := new(types.Project)
	err := s.db.Get(dst, projectSelectToken, token)
	return dst, err
}

// List returns a list of projects by user.
func (s *ProjectStore) List(ctx context.Context, user int64, opts types.Params) ([]*types.Project, error) {
	dst := []*types.Project{}
	err := s.db.Select(&dst, projectSelect, user)
	// TODO(bradrydzewski) add limit and offset
	return dst, err
}

// Create saves the project details.
func (s *ProjectStore) Create(ctx context.Context, project *types.Project) error {
	query := projectInsert

	if s.db.DriverName() == "postgres" {
		query = projectInsertPg
	}

	query, arg, err := s.db.BindNamed(query, project)
	if err != nil {
		return err
	}

	if s.db.DriverName() == "postgres" {
		return s.db.QueryRow(query, arg...).Scan(&project.ID)
	}

	res, err := s.db.Exec(query, arg...)
	if err != nil {
		return err
	}
	project.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// Update updates the project details.
func (s *ProjectStore) Update(ctx context.Context, project *types.Project) error {
	query, arg, err := s.db.BindNamed(projectUpdate, project)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, arg...)
	return err
}

// Delete deletes the project.
func (s *ProjectStore) Delete(ctx context.Context, project *types.Project) error {
	_, err := s.db.Exec(projectDelete, project.ID)
	return err
}

const projectBase = `
SELECT
 project_id
,project_name
,project_desc
,project_token
,project_active
,project_created
,project_updated
FROM projects
`

const projectSelect = projectBase + `
WHERE project_id IN (
  SELECT member_project_id
  FROM members
  WHERE member_user_id = $1
)
ORDER BY project_name
`

const projectSelectID = projectBase + `
WHERE project_id = $1
`

const projectSelectToken = projectBase + `
WHERE project_token = $1
`

const projectDelete = `
DELETE FROM projects
WHERE project_id = $1
`

const projectInsert = `
INSERT INTO projects (
 project_name
,project_desc
,project_token
,project_active
,project_created
,project_updated
) values (
 :project_name
,:project_desc
,:project_token
,:project_active
,:project_created
,:project_updated
)
`

const projectInsertPg = projectInsert + `
RETURNING project_id
`

const projectUpdate = `
UPDATE projects
SET
 project_name      = :project_name
,project_desc      = :project_desc
,project_active    = :project_active
,project_updated   = :project_updated
WHERE project_id = :project_id
`

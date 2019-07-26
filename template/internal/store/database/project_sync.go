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
	"github.com/{{github}}/internal/store/database/mutex"
	"github.com/{{github}}/types"
)

var _ store.ProjectStore = (*ProjectStoreSync)(nil)

// NewProjectStoreSync returns a new ProjectStoreSync.
func NewProjectStoreSync(store *ProjectStore) *ProjectStoreSync {
	return &ProjectStoreSync{store}
}

// ProjectStoreSync synronizes read and write access to the
// project store. This prevents race conditions when the database
// type is sqlite3.
type ProjectStoreSync struct{ *ProjectStore }

// Find finds the project by id.
func (s *ProjectStoreSync) Find(ctx context.Context, id int64) (*types.Project, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.ProjectStore.Find(ctx, id)
}

// List returns a list of projects by user.
func (s *ProjectStoreSync) List(ctx context.Context, id int64, opts types.Params) ([]*types.Project, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.ProjectStore.List(ctx, id, opts)
}

// Create saves the project details.
func (s *ProjectStoreSync) Create(ctx context.Context, project *types.Project) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.ProjectStore.Create(ctx, project)
}

// Update updates the project details.
func (s *ProjectStoreSync) Update(ctx context.Context, project *types.Project) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.ProjectStore.Update(ctx, project)
}

// Delete deletes the project.
func (s *ProjectStoreSync) Delete(ctx context.Context, project *types.Project) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.ProjectStore.Delete(ctx, project)
}

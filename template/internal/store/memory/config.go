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

// Package memory provides readonly memory data storage.
package memory

import (
	"context"

	"github.com/{{github}}/types"
)

// New returns a new system configuration store.
func New(config *types.Config) *SystemStore {
	return &SystemStore{config: config}
}

// SystemStore is a system store that loads system
// configuration parameters stored in the environment.
type SystemStore struct {
	config *types.Config
}

// Config returns the system configuration.
func (c *SystemStore) Config(ctx context.Context) *types.Config {
	return c.config
}

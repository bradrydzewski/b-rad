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

package server

import (
	"os"

	"github.com/{{github}}/types"

	"github.com/kelseyhightower/envconfig"
)

// legacy environment variables. the key is the legacy
// variable name, and the value is the new variable name.
var legacy = map[string]string{
	// none defined
}

// load returns the system configuration from the
// host environment.
func load() (*types.Config, error) {

	// loop through legacy environment variable and, if set
	// rewrite to the new variable name.
	for k, v := range legacy {
		if s, ok := os.LookupEnv(k); ok {
			os.Setenv(v, s)
		}
	}

	config := new(types.Config)
	// read the configuration from the environment and
	// populate the configuration structure.
	err := envconfig.Process("", config)
	return config, err
}

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

package cli

import (
	"os"

	"github.com/{{github}}/cli/util"

	"gopkg.in/alecthomas/kingpin.v2"
)

type logoutCommand struct{}

func (c *logoutCommand) run(*kingpin.ParseContext) error {
	path, err := util.Config()
	if err != nil {
		return err
	}
	return os.Remove(path)
}

// helper function to register the logout command.
func registerLogout(app *kingpin.Application) {
	c := new(logoutCommand)

	app.Command("logout", "logout from the remote server").
		Action(c.run)
}

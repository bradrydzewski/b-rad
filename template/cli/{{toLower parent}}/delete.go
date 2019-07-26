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

package {{toLower parent}}

import (
	"github.com/{{github}}/cli/util"

	"gopkg.in/alecthomas/kingpin.v2"
)

type deleteCommand struct {
	proj int64
	id   int64
}

func (c *deleteCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	return client.{{parent}}Delete(c.proj, c.id)
}

// helper function registers the user delete command
func registerDelete(app *kingpin.CmdClause) {
	c := new(deleteCommand)

	cmd := app.Command("delete", "delete a {{toLower parent}}").
		Action(c.run)

	cmd.Arg("project_id", "project id").
		Required().
		Int64Var(&c.proj)

	cmd.Arg("{{toLower parent}}_id ", "{{toLower parent}} id").
		Required().
		Int64Var(&c.id)
}

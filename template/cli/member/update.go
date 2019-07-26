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

package member

import (
	"os"
	"text/template"

	"github.com/{{github}}/cli/util"
	"github.com/{{github}}/types"
	"github.com/{{github}}/types/enum"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
)

type updateCommand struct {
	proj int64
	user string
	role string
	tmpl string
}

func (c *updateCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}

	in := new(types.MembershipInput)
	in.Project = c.proj
	in.User = c.user
	switch c.role {
	case "admin":
		in.Role = enum.RoleAdmin
	case "developer":
		in.Role = enum.RoleDeveloper
	}

	member, err := client.MemberUpdate(in)
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, member)
}

// helper function registers the user update command
func registerUpdate(app *kingpin.CmdClause) {
	c := new(updateCommand)

	cmd := app.Command("update", "update a project").
		Action(c.run)

	cmd.Arg("project", "project id").
		Required().
		Int64Var(&c.proj)

	cmd.Arg("user id or email", "member id or email").
		Required().
		StringVar(&c.user)

	cmd.Flag("role", "update member role").
		StringVar(&c.role)

	cmd.Flag("format", "format the output using a Go template").
		Default(memberTmpl).
		Hidden().
		StringVar(&c.tmpl)
}

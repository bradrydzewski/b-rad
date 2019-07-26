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

package project

import (
	"os"
	"text/template"

	"github.com/{{github}}/cli/util"
	"github.com/{{github}}/types"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
)

type createCommand struct {
	name string
	desc string
	tmpl string
}

func (c *createCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	in := &types.Project{
		Name: c.name,
		Desc: c.desc,
	}
	proj, err := client.ProjectCreate(in)
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, proj)
}

// helper function registers the user create command
func registerCreate(app *kingpin.CmdClause) {
	c := new(createCommand)

	cmd := app.Command("create", "create a project").
		Action(c.run)

	cmd.Arg("name", "project name").
		Required().
		StringVar(&c.name)

	cmd.Flag("desc", "project description").
		StringVar(&c.desc)

	cmd.Flag("format", "format the output using a Go template").
		Default(projectTmpl).
		Hidden().
		StringVar(&c.tmpl)
}

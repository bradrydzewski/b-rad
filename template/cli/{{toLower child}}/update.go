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

package {{toLower child}}

import (
	"os"
	"text/template"

	"github.com/{{github}}/cli/util"
	"github.com/{{github}}/types"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/guregu/null.v4"
)

type updateCommand struct {
	proj int64
	{{toLower parent}}  int64
	{{toLower child}}  int64
	name string
	desc string
	tmpl string
}

func (c *updateCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}

	in := new(types.{{child}}Input)
	if v := c.name; v != "" {
		in.Name = null.StringFrom(v)
	}
	if v := c.desc; v != "" {
		in.Desc = null.StringFrom(v)
	}

	{{toLower parent}}, err := client.{{child}}Update(c.proj, c.{{toLower parent}}, c.{{toLower child}}, in)
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, {{toLower parent}})
}

// helper function registers the update command
func registerUpdate(app *kingpin.CmdClause) {
	c := new(updateCommand)

	cmd := app.Command("update", "update a {{toLower parent}}").
		Action(c.run)

	cmd.Arg("project_id", "project id").
		Required().
		Int64Var(&c.proj)

	cmd.Arg("{{toLower parent}}_id", "{{toLower parent}} id").
		Required().
		Int64Var(&c.{{toLower parent}})

	cmd.Arg("{{toLower child}}_id", "{{toLower child}} id").
		Required().
		Int64Var(&c.{{toLower child}})

	cmd.Flag("name", "update project name").
		StringVar(&c.name)

	cmd.Flag("desc", "update project description").
		StringVar(&c.desc)

	cmd.Flag("format", "format the output using a Go template").
		Default(projectTmpl).
		Hidden().
		StringVar(&c.tmpl)
}

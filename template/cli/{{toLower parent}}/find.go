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
	"os"
	"text/template"

	"github.com/{{github}}/cli/util"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
)

type findCommand struct {
	proj int64
	id   int64
	tmpl string
}

func (c *findCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	proj, err := client.{{parent}}(c.proj, c.id)
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl + "\n")
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, proj)
}

// helper function registers the user find command
func registerFind(app *kingpin.CmdClause) {
	c := new(findCommand)

	cmd := app.Command("find", "display project details").
		Action(c.run)

	cmd.Arg("project_id", "project id").
		Required().
		Int64Var(&c.proj)

	cmd.Arg("{{toLower parent}}_id", "{{toLower parent}} id").
		Required().
		Int64Var(&c.id)

	cmd.Flag("format", "format the output using a Go template").
		Default(projectTmpl).
		Hidden().
		StringVar(&c.tmpl)
}

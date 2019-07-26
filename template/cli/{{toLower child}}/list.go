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
	"github.com/drone/funcmap"

	"gopkg.in/alecthomas/kingpin.v2"
)

const projectTmpl = `
id:   {{`{{ .ID }}`}}
name: {{`{{ .Name }}`}}
desc: {{`{{ .Desc }}`}}
`

type listCommand struct {
	proj int64
	{{toLower parent}}  int64
	tmpl string
}

func (c *listCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	list, err := client.{{child}}List(c.proj, c.{{toLower parent}})
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl + "\n")
	if err != nil {
		return err
	}
	for _, item := range list {
		tmpl.Execute(os.Stdout, item)
	}
	return nil
}

// helper function registers the list command
func registerList(app *kingpin.CmdClause) {
	c := new(listCommand)

	cmd := app.Command("ls", "display a list of {{toLower parent}}s").
		Action(c.run)

	cmd.Arg("project_id", "project id").
		Required().
		Int64Var(&c.proj)

	cmd.Arg("{{toLower parent}}_id", "{{toLower parent}} id").
		Required().
		Int64Var(&c.{{toLower parent}})

	cmd.Flag("format", "format the output using a Go template").
		Default(projectTmpl).
		Hidden().
		StringVar(&c.tmpl)
}

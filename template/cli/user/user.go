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

package user

import (
	"os"
	"text/template"

	"github.com/{{github}}/cli/util"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
)

const userTmpl = `
email: {{`{{ .Email }}`}}
admin: {{`{{ .Admin }}`}}
`

type command struct {
	tmpl string
}

func (c *command) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	user, err := client.Self()
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, user)
}

// Register the command.
func Register(app *kingpin.Application) {
	c := new(command)

	cmd := app.Command("account", "display authenticated user").
		Action(c.run)

	cmd.Flag("format", "format the output using a Go template").
		Default(userTmpl).
		Hidden().
		StringVar(&c.tmpl)
}

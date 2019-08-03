// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower parent}}

import (
	"github.com/{{toLower repo}}/cli/util"

	"gopkg.in/alecthomas/kingpin.v2"
)

type deleteCommand struct {
	{{toLower project}} string
	slug    string
}

func (c *deleteCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	return client.{{title parent}}Delete(c.{{toLower project}}, c.slug)
}

// helper function registers the user delete command
func registerDelete(app *kingpin.CmdClause) {
	c := new(deleteCommand)

	cmd := app.Command("delete", "delete a {{toLower parent}}").
		Action(c.run)

	cmd.Arg("{{toLower project}} ", "{{toLower project}} slug").
		Required().
		StringVar(&c.{{toLower project}})

	cmd.Arg("slug ", "{{toLower parent}} slug").
		Required().
		StringVar(&c.slug)
}

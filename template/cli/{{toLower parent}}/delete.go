// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower parent}}

import (
	"github.com/{{toLower repo}}/cli/util"

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
	return client.{{title parent}}Delete(c.proj, c.id)
}

// helper function registers the user delete command
func registerDelete(app *kingpin.CmdClause) {
	c := new(deleteCommand)

	cmd := app.Command("delete", "delete a {{toLower parent}}").
		Action(c.run)

	cmd.Arg("{{toLower project}}_id", "{{toLower project}} id").
		Required().
		Int64Var(&c.proj)

	cmd.Arg("{{toLower parent}}_id ", "{{toLower parent}} id").
		Required().
		Int64Var(&c.id)
}

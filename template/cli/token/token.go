// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package token

import (
	"encoding/json"
	"os"

	"github.com/{{toLower repo}}/cli/util"

	"gopkg.in/alecthomas/kingpin.v2"
)

type command struct {
	json bool
}

func (c *command) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	token, err := client.Token()
	if err != nil {
		return err
	}
	if c.json {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(token)
	}
	println(token.Value)
	return nil
}

// Register the command.
func Register(app *kingpin.Application) {
	c := new(command)

	cmd := app.Command("token", "generate a personal token").
		Action(c.run)

	cmd.Flag("json", "json encode the output").
		BoolVar(&c.json)

}

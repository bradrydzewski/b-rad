// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package users

import (
	"encoding/json"
	"os"
	"text/template"

	"github.com/{{toLower repo}}/cli/util"
	"github.com/{{toLower repo}}/types"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
)

type createCommand struct {
	email string
	admin bool
	tmpl  string
	json  bool
}

func (c *createCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	in := &types.User{
		Admin:    c.admin,
		Email:    c.email,
		Password: util.Password(),
	}
	user, err := client.UserCreate(in)
	if err != nil {
		return err
	}
	if c.json {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(user)
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, user)
}

// helper function registers the user create command
func registerCreate(app *kingpin.CmdClause) {
	c := new(createCommand)

	cmd := app.Command("create", "create a user").
		Action(c.run)

	cmd.Arg("email", "user email").
		Required().
		StringVar(&c.email)

	cmd.Arg("admin", "user is admin").
		BoolVar(&c.admin)

	cmd.Flag("json", "json encode the output").
		BoolVar(&c.json)

	cmd.Flag("format", "format the output using a Go template").
		Default(userTmpl).
		Hidden().
		StringVar(&c.tmpl)
}

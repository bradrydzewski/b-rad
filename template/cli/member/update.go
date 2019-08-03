// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package member

import (
	"encoding/json"
	"os"
	"text/template"

	"github.com/{{toLower repo}}/cli/util"
	"github.com/{{toLower repo}}/types"
	"github.com/{{toLower repo}}/types/enum"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
)

type updateCommand struct {
	slug string
	user string
	role string
	tmpl string
	json bool
}

func (c *updateCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}

	in := new(types.MembershipInput)
	in.{{title project}} = c.slug
	in.User = c.user
	switch c.role {
	case "admin":
		in.Role = enum.RoleAdmin
	case "developer":
		in.Role = enum.RoleDeveloper
	}

	item, err := client.MemberUpdate(in)
	if err != nil {
		return err
	}
	if c.json {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(item)
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, item)
}

// helper function registers the user update command
func registerUpdate(app *kingpin.CmdClause) {
	c := new(updateCommand)

	cmd := app.Command("update", "update a {{toLower project}}").
		Action(c.run)

	cmd.Arg("{{toLower project}}", "{{toLower project}} slug").
		Required().
		StringVar(&c.slug)

	cmd.Arg("user id or email", "member id or email").
		Required().
		StringVar(&c.user)

	cmd.Flag("role", "update member role").
		StringVar(&c.role)

	cmd.Flag("json", "json encode the output").
		BoolVar(&c.json)

	cmd.Flag("format", "format the output using a Go template").
		Default(memberTmpl).
		Hidden().
		StringVar(&c.tmpl)
}

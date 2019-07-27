// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower parent}}

import (
	"os"
	"text/template"

	"github.com/{{toLower repo}}/cli/util"
	"github.com/{{toLower repo}}/types"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/guregu/null.v4"
)

type updateCommand struct {
	proj int64
	id   int64
	name string
	desc string
	tmpl string
}

func (c *updateCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}

	in := new(types.{{title parent}}Input)
	if v := c.name; v != "" {
		in.Name = null.StringFrom(v)
	}
	if v := c.desc; v != "" {
		in.Desc = null.StringFrom(v)
	}

	{{toLower parent}}, err := client.{{title parent}}Update(c.proj, c.id, in)
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

	cmd.Arg("{{toLower project}}_id", "{{toLower project}} id").
		Required().
		Int64Var(&c.proj)

	cmd.Arg("{{toLower parent}}_id", "{{toLower parent}} id").
		Required().
		Int64Var(&c.id)

	cmd.Flag("name", "update {{toLower project}} name").
		StringVar(&c.name)

	cmd.Flag("desc", "update {{toLower project}} description").
		StringVar(&c.desc)

	cmd.Flag("format", "format the output using a Go template").
		Default({{toLower project}}Tmpl).
		Hidden().
		StringVar(&c.tmpl)
}

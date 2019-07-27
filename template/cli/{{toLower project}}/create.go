// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower project}}

import (
	"os"
	"text/template"

	"github.com/{{toLower repo}}/cli/util"
	"github.com/{{toLower repo}}/types"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
)

type createCommand struct {
	name string
	desc string
	tmpl string
}

func (c *createCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	in := &types.{{title project}}{
		Name: c.name,
		Desc: c.desc,
	}
	proj, err := client.{{title project}}Create(in)
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, proj)
}

// helper function registers the user create command
func registerCreate(app *kingpin.CmdClause) {
	c := new(createCommand)

	cmd := app.Command("create", "create a {{toLower project}}").
		Action(c.run)

	cmd.Arg("name", "{{toLower project}} name").
		Required().
		StringVar(&c.name)

	cmd.Flag("desc", "{{toLower project}} description").
		StringVar(&c.desc)

	cmd.Flag("format", "format the output using a Go template").
		Default({{toLower project}}Tmpl).
		Hidden().
		StringVar(&c.tmpl)
}

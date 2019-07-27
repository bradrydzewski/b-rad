// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower child}}

import (
	"os"
	"text/template"

	"github.com/{{toLower repo}}/cli/util"
	"github.com/{{toLower repo}}/types"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
)

type createCommand struct {
	proj int64
	{{toLower parent}}  int64
	name string
	desc string
	tmpl string
}

func (c *createCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	in := &types.{{title child}}{
		Name: c.name,
		Desc: c.desc,
	}
	proj, err := client.{{title child}}Create(c.proj, c.{{toLower parent}}, in)
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

	cmd := app.Command("create", "create a {{toLower child}}").
		Action(c.run)

	cmd.Arg("{{toLower project}}_id", "{{toLower project}} id").
		Required().
		Int64Var(&c.proj)

	cmd.Arg("{{toLower parent}}_id", "{{toLower parent}} id").
		Required().
		Int64Var(&c.{{toLower parent}})

	cmd.Arg("name", "{{toLower parent}} name").
		Required().
		StringVar(&c.name)

	cmd.Flag("desc", "{{toLower parent}} description").
		StringVar(&c.desc)

	cmd.Flag("format", "format the output using a Go template").
		Default({{toLower project}}Tmpl).
		Hidden().
		StringVar(&c.tmpl)
}

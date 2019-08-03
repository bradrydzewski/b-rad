// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower parent}}

import (
	"encoding/json"
	"os"
	"text/template"

	"github.com/{{toLower repo}}/cli/util"
	"github.com/{{toLower repo}}/types"
	"github.com/drone/funcmap"

	"gopkg.in/alecthomas/kingpin.v2"
)

const {{toLower parent}}Tmpl = `
id:   {{`{{`}} .ID {{`}}`}}
slug: {{`{{`}} .Slug {{`}}`}}
name: {{`{{`}} .Name {{`}}`}}
desc: {{`{{`}} .Desc {{`}}`}}
`

type listCommand struct {
	slug string
	tmpl string
	json bool
	page int
	size int
}

func (c *listCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	list, err := client.{{title parent}}List(c.slug, types.Params{
		Size: c.size,
		Page: c.page,
	})
	if err != nil {
		return err
	}
	if c.json {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(list)
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

	cmd.Arg("{{toLower project}} ", "{{toLower project}} slug").
		Required().
		StringVar(&c.slug)

	cmd.Flag("page", "page number").
		IntVar(&c.page)

	cmd.Flag("per-page", "page size").
		IntVar(&c.size)

	cmd.Flag("json", "json encode the output").
		BoolVar(&c.json)

	cmd.Flag("format", "format the output using a Go template").
		Default({{toLower parent}}Tmpl).
		Hidden().
		StringVar(&c.tmpl)
}

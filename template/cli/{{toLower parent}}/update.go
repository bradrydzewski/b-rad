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
	"github.com/gotidy/ptr"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
)

type updateCommand struct {
	{{toLower project}} string
	slug    string
	name    string
	desc    string
	tmpl    string
	json    bool
}

func (c *updateCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}

	in := new(types.{{title parent}}Input)
	if v := c.name; v != "" {
		in.Name = ptr.String(v)
	}
	if v := c.desc; v != "" {
		in.Desc = ptr.String(v)
	}

	item, err := client.{{title parent}}Update(c.{{toLower project}}, c.slug, in)
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

// helper function registers the update command
func registerUpdate(app *kingpin.CmdClause) {
	c := new(updateCommand)

	cmd := app.Command("update", "update a {{toLower parent}}").
		Action(c.run)

	cmd.Arg("{{toLower project}} ", "{{toLower project}} slug").
		Required().
		StringVar(&c.{{toLower project}})

	cmd.Arg("slug ", "{{toLower parent}} slug").
		Required().
		StringVar(&c.slug)

	cmd.Flag("name", "update {{toLower project}} name").
		StringVar(&c.name)

	cmd.Flag("desc", "update {{toLower project}} description").
		StringVar(&c.desc)

	cmd.Flag("json", "json encode the output").
		BoolVar(&c.json)

	cmd.Flag("format", "format the output using a Go template").
		Default({{toLower parent}}Tmpl).
		Hidden().
		StringVar(&c.tmpl)
}

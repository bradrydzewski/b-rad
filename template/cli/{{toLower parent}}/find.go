// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower parent}}

import (
	"encoding/json"
	"os"
	"text/template"

	"github.com/{{toLower repo}}/cli/util"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
)

type findCommand struct {
	{{toLower project}} string
	slug    string
	tmpl    string
	json    bool
}

func (c *findCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	item, err := client.{{title parent}}(c.{{toLower project}}, c.slug)
	if err != nil {
		return err
	}
	if c.json {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(item)
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl + "\n")
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, item)
}

// helper function registers the user find command
func registerFind(app *kingpin.CmdClause) {
	c := new(findCommand)

	cmd := app.Command("find", "display {{toLower project}} details").
		Action(c.run)

	cmd.Arg("{{toLower project}} ", "{{toLower project}} slug").
		Required().
		StringVar(&c.{{toLower project}})

	cmd.Arg("slug ", "{{toLower parent}} slug").
		Required().
		StringVar(&c.slug)

	cmd.Flag("json", "json encode the output").
		BoolVar(&c.json)

	cmd.Flag("format", "format the output using a Go template").
		Default({{toLower parent}}Tmpl).
		Hidden().
		StringVar(&c.tmpl)
}

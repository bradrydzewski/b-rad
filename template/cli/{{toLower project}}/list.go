// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower project}}

import (
	"os"
	"text/template"

	"github.com/{{toLower repo}}/cli/util"
	"github.com/drone/funcmap"

	"gopkg.in/alecthomas/kingpin.v2"
)

const {{toLower project}}Tmpl = `
id:   {{`{{`}} .ID {{`}}`}}
name: {{`{{`}} .Name {{`}}`}}
desc: {{`{{`}} .Desc {{`}}`}}
`

type listCommand struct {
	tmpl string
}

func (c *listCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	list, err := client.{{title project}}List()
	if err != nil {
		return err
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

// helper function registers the user list command
func registerList(app *kingpin.CmdClause) {
	c := new(listCommand)

	cmd := app.Command("ls", "display a list of {{toLower project}}s").
		Action(c.run)

	cmd.Flag("format", "format the output using a Go template").
		Default({{toLower project}}Tmpl).
		Hidden().
		StringVar(&c.tmpl)
}

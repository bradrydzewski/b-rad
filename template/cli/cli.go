// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package cli

import (
	"context"
	"os"

	"github.com/{{toLower repo}}/cli/{{toLower child}}"
	"github.com/{{toLower repo}}/cli/{{toLower parent}}"
	"github.com/{{toLower repo}}/cli/member"
	"github.com/{{toLower repo}}/cli/{{toLower project}}"
	"github.com/{{toLower repo}}/cli/server"
	"github.com/{{toLower repo}}/cli/token"
	"github.com/{{toLower repo}}/cli/user"
	"github.com/{{toLower repo}}/cli/users"
	"github.com/{{toLower repo}}/version"

	"gopkg.in/alecthomas/kingpin.v2"
)

// empty context
var nocontext = context.Background()

// application name
var application = "{{toLower name}}"

// application description
var description = "description goes here" // TODO edit this application description

// Command parses the command line arguments and then executes a
// subcommand program.
func Command() {
	app := kingpin.New(application, description)
	server.Register(app)
	user.Register(app)
	{{toLower project}}.Register(app)
	{{toLower parent}}.Register(app)
	{{toLower child}}.Register(app)
	member.Register(app)
	users.Register(app)
	token.Register(app)
	registerLogin(app)
	registerLogout(app)
	registerRegister(app)

	kingpin.Version(version.Version.String())
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

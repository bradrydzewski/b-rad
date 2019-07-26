// Copyright 2019 Brad Rydzewski. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cli

import (
	"context"
	"os"

	"github.com/{{github}}/cli/{{toLower child}}"
	"github.com/{{github}}/cli/{{toLower parent}}"
	"github.com/{{github}}/cli/member"
	"github.com/{{github}}/cli/project"
	"github.com/{{github}}/cli/server"
	"github.com/{{github}}/cli/token"
	"github.com/{{github}}/cli/user"
	"github.com/{{github}}/cli/users"
	"github.com/{{github}}/version"

	"gopkg.in/alecthomas/kingpin.v2"
)

// empty context
var nocontext = context.Background()

// application name
var application = "{{app}}"

// application description
var description = "description goes here" // TODO edit this application description

// Command parses the command line arguments and then executes a
// subcommand program.
func Command() {
	app := kingpin.New(application, description)
	server.Register(app)
	user.Register(app)
	project.Register(app)
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

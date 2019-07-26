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
	"encoding/json"
	"io/ioutil"

	"github.com/{{github}}/cli/util"
	"github.com/{{github}}/client"

	"gopkg.in/alecthomas/kingpin.v2"
)

type loginCommand struct {
	server string
}

func (c *loginCommand) run(*kingpin.ParseContext) error {
	username, password := util.Credentials()
	client := client.New(c.server)
	token, err := client.Login(username, password)
	if err != nil {
		return err
	}
	path, err := util.Config()
	if err != nil {
		return err
	}
	token.Address = c.server
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0600)
}

// helper function to register the logout command.
func registerLogin(app *kingpin.Application) {
	c := new(loginCommand)

	cmd := app.Command("login", "login to the remote server").
		Action(c.run)

	cmd.Arg("server", "server address").
		Default("http://localhost:3000").
		StringVar(&c.server)
}

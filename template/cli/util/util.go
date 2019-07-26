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

package util

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/adrg/xdg"
	"github.com/{{github}}/client"
	"github.com/{{github}}/types"
	"golang.org/x/crypto/ssh/terminal"
)

// Client returns a client that is configured from file.
func Client() (*client.HTTPClient, error) {
	path, err := Config()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	token := new(types.Token)
	if err := json.Unmarshal(data, token); err != nil {
		return nil, err
	}
	if time.Now().Unix() > token.Expires.Unix() {
		return nil, errors.New("token is expired, please login")
	}
	client := client.NewToken(token.Address, token.Value)
	if os.Getenv("DEBUG") == "true" {
		client.SetDebug(true)
	}
	return client, nil
}

// Config returns the configuration file path.
func Config() (string, error) {
	return xdg.ConfigFile(
		filepath.Join("app", "config.json"),
	)
}

// Credentials returns the username and password from stdin.
func Credentials() (string, string) {
	return Username(), Password()
}

// Username returns the username from stdin.
func Username() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	return strings.TrimSpace(username)
}

// Password returns the password from stdin.
func Password() string {
	fmt.Print("Enter Password: ")
	passwordb, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(passwordb)

	return strings.TrimSpace(password)
}

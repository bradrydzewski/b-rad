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

package database

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var noContext = context.Background()

// connect opens a new test database connection.
func connect() (*sqlx.DB, error) {
	var (
		driver = "sqlite3"
		config = ":memory:?_foreign_keys=1"
	)
	if os.Getenv("DATABASE_CONFIG") != "" {
		driver = os.Getenv("DATABASE_DRIVER")
		config = os.Getenv("DATABASE_CONFIG")
	}
	return Connect(driver, config)
}

// seed seed the database state.
func seed(db *sqlx.DB) error {
	db.Exec("TRUNCATE TABLE {{toLower child}}s")
	db.Exec("TRUNCATE TABLE {{toLower parent}}s")
	db.Exec("TRUNCATE TABLE members")
	db.Exec("TRUNCATE TABLE projects")
	db.Exec("TRUNCATE TABLE users")

	out, err := ioutil.ReadFile("testdata/seed.sql")
	if err != nil {
		return err
	}
	parts := strings.Split(string(out), ";")
	for _, stmt := range parts {
		if stmt == "" {
			continue
		}
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("%s: %s", err, stmt)
		}
	}
	return nil
}

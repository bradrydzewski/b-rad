package sqlite

import (
	"database/sql"
)

var migrations = []struct {
	name string
	stmt string
}{
	{
		name: "create-table-users",
		stmt: createTableUsers,
	},
	{
		name: "create-table-{{toLower project}}s",
		stmt: createTable{{title project}}s,
	},
	{
		name: "create-table-members",
		stmt: createTableMembers,
	},
	{
		name: "create-index-members-{{toLower project}}-id",
		stmt: createIndexMembers{{title project}}Id,
	},
	{
		name: "create-index-members-user-id",
		stmt: createIndexMembersUserId,
	},
	{
		name: "create-table-{{toLower parent}}s",
		stmt: createTable{{title parent}}s,
	},
	{
		name: "create-index-{{toLower parent}}-{{toLower project}}-id",
		stmt: createIndex{{title parent}}{{title project}}Id,
	},
	{
		name: "create-table-{{toLower child}}s",
		stmt: createTable{{title child}}s,
	},
	{
		name: "create-index-{{toLower child}}-{{toLower parent}}-id",
		stmt: createIndex{{title child}}{{title parent}}Id,
	},
}

// Migrate performs the database migration. If the migration fails
// and error is returned.
func Migrate(db *sql.DB) error {
	if err := createTable(db); err != nil {
		return err
	}
	completed, err := selectCompleted(db)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	for _, migration := range migrations {
		if _, ok := completed[migration.name]; ok {

			continue
		}

		if _, err := db.Exec(migration.stmt); err != nil {
			return err
		}
		if err := insertMigration(db, migration.name); err != nil {
			return err
		}

	}
	return nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(migrationTableCreate)
	return err
}

func insertMigration(db *sql.DB, name string) error {
	_, err := db.Exec(migrationInsert, name)
	return err
}

func selectCompleted(db *sql.DB) (map[string]struct{}, error) {
	migrations := map[string]struct{}{}
	rows, err := db.Query(migrationSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		migrations[name] = struct{}{}
	}
	return migrations, nil
}

//
// migration table ddl and sql
//

var migrationTableCreate = `
CREATE TABLE IF NOT EXISTS migrations (
 name VARCHAR(255)
,UNIQUE(name)
)
`

var migrationInsert = `
INSERT INTO migrations (name) VALUES (?)
`

var migrationSelect = `
SELECT name FROM migrations
`

//
// 001_create_table_user.sql
//

var createTableUsers = `
CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT
,user_email         TEXT COLLATE NOCASE
,user_password      TEXT
,user_token         TEXT
,user_name          TEXT
,user_company       TEXT
,user_admin         BOOLEAN
,user_blocked       BOOLEAN
,user_created       INTEGER
,user_updated       INTEGER
,user_authed        INTEGER
,UNIQUE(user_token)
,UNIQUE(user_email COLLATE NOCASE)
);
`

//
// 002_create_table_{{toLower project}}.sql
//

var createTable{{title project}}s = `
CREATE TABLE IF NOT EXISTS {{toLower project}}s (
 {{toLower project}}_id          INTEGER PRIMARY KEY AUTOINCREMENT
,{{toLower project}}_name        TEXT
,{{toLower project}}_desc        TEXT
,{{toLower project}}_token       TEXT
,{{toLower project}}_active      BOOLEAN
,{{toLower project}}_created     INTEGER
,{{toLower project}}_updated     INTEGER
,UNIQUE({{toLower project}}_token)
);
`

//
// 003_create_table_member.sql
//

var createTableMembers = `
CREATE TABLE IF NOT EXISTS members (
 member_{{toLower project}}_id INTEGER
,member_user_id    INTEGER
,member_role       INTEGER
,PRIMARY KEY(member_{{toLower project}}_id, member_user_id)
);
`

var createIndexMembers{{title project}}Id = `
CREATE INDEX IF NOT EXISTS index_members_{{toLower project}} ON members(member_{{toLower project}}_id)
`

var createIndexMembersUserId = `
CREATE INDEX IF NOT EXISTS index_members_user ON members(member_user_id)
`

//
// 005_create_table_{{toLower parent}}.sql
//

var createTable{{title parent}}s = `
CREATE TABLE IF NOT EXISTS {{toLower parent}}s (
 {{toLower parent}}_id          INTEGER PRIMARY KEY AUTOINCREMENT
,{{toLower parent}}_{{toLower project}}_id  INTEGER
,{{toLower parent}}_name        TEXT
,{{toLower parent}}_desc        TEXT
,{{toLower parent}}_created     INTEGER
,{{toLower parent}}_updated     INTEGER
);
`

var createIndex{{title parent}}{{title project}}Id = `
CREATE INDEX IF NOT EXISTS index_{{toLower parent}}_{{toLower project}} ON {{toLower parent}}s({{toLower parent}}_{{toLower project}}_id);
`

//
// 007_create_table_{{toLower child}}.sql
//

var createTable{{title child}}s = `
CREATE TABLE IF NOT EXISTS {{toLower child}}s (
 {{toLower child}}_id       INTEGER PRIMARY KEY AUTOINCREMENT
,{{toLower child}}_{{toLower parent}}_id   INTEGER
,{{toLower child}}_name     TEXT
,{{toLower child}}_desc     TEXT
,{{toLower child}}_created  INTEGER
,{{toLower child}}_updated  INTEGER
);
`

var createIndex{{title child}}{{title parent}}Id = `
CREATE INDEX IF NOT EXISTS index_{{toLower child}}_{{toLower parent}} ON {{toLower child}}s({{toLower child}}_{{toLower parent}}_id);
`

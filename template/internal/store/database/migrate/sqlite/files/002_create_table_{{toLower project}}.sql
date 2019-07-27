-- name: create-table-{{toLower project}}s

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

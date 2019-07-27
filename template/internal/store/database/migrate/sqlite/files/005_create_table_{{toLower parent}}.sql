-- name: create-table-{{toLower parent}}s

CREATE TABLE IF NOT EXISTS {{toLower parent}}s (
 {{toLower parent}}_id          INTEGER PRIMARY KEY AUTOINCREMENT
,{{toLower parent}}_{{toLower project}}_id  INTEGER
,{{toLower parent}}_name        TEXT
,{{toLower parent}}_desc        TEXT
,{{toLower parent}}_created     INTEGER
,{{toLower parent}}_updated     INTEGER
);

-- name: create-index-{{toLower parent}}-{{toLower project}}-id

CREATE INDEX IF NOT EXISTS index_{{toLower parent}}_{{toLower project}} ON {{toLower parent}}s({{toLower parent}}_{{toLower project}}_id);

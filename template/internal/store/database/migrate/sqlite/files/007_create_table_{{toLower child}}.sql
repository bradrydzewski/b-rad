-- name: create-table-{{toLower child}}s

CREATE TABLE IF NOT EXISTS {{toLower child}}s (
 {{toLower child}}_id       INTEGER PRIMARY KEY AUTOINCREMENT
,{{toLower child}}_{{toLower parent}}_id   INTEGER
,{{toLower child}}_name     TEXT
,{{toLower child}}_desc     TEXT
,{{toLower child}}_created  INTEGER
,{{toLower child}}_updated  INTEGER
);

-- name: create-index-{{toLower child}}-{{toLower parent}}-id

CREATE INDEX IF NOT EXISTS index_{{toLower child}}_{{toLower parent}} ON {{toLower child}}s({{toLower child}}_{{toLower parent}}_id);

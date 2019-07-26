-- name: create-table-{{toLower child}}s

CREATE TABLE IF NOT EXISTS {{toLower child}}s (
 {{toLower child}}_id       SERIAL PRIMARY KEY
,{{toLower child}}_{{toLower parent}}_id   INTEGER
,{{toLower child}}_name     VARCHAR(250)
,{{toLower child}}_desc     VARCHAR(2000)
,{{toLower child}}_created  INTEGER
,{{toLower child}}_updated  INTEGER
);

-- name: create-index-{{toLower child}}-{{toLower parent}}-id

CREATE INDEX IF NOT EXISTS index_{{toLower child}}_{{toLower parent}} ON {{toLower child}}s({{toLower child}}_{{toLower parent}}_id);

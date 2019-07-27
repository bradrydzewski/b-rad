-- name: create-table-{{toLower project}}s

CREATE TABLE IF NOT EXISTS {{toLower project}}s (
 {{toLower project}}_id          SERIAL PRIMARY KEY
,{{toLower project}}_name        VARCHAR(250)
,{{toLower project}}_desc        VARCHAR(250)
,{{toLower project}}_token       VARCHAR(250)
,{{toLower project}}_active      BOOLEAN
,{{toLower project}}_created     INTEGER
,{{toLower project}}_updated     INTEGER
,UNIQUE({{toLower project}}_token)
);

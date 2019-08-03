CREATE TABLE IF NOT EXISTS {{toLower parent}}s (
 {{toLower parent}}_id          INTEGER PRIMARY KEY AUTOINCREMENT
,{{toLower parent}}_{{toLower project}}_id  INTEGER
,{{toLower parent}}_slug        TEXT
,{{toLower parent}}_name        TEXT
,{{toLower parent}}_desc        TEXT
,{{toLower parent}}_created     INTEGER
,{{toLower parent}}_updated     INTEGER
,UNIQUE({{toLower parent}}_{{toLower project}}_id, {{toLower parent}}_slug)
);

CREATE TABLE IF NOT EXISTS {{toLower child}}s (
 {{toLower child}}_id         INTEGER PRIMARY KEY AUTOINCREMENT
,{{toLower child}}_{{toLower project}}_id INTEGER
,{{toLower child}}_{{toLower parent}}_id   INTEGER
,{{toLower child}}_slug       TEXT
,{{toLower child}}_name       TEXT
,{{toLower child}}_desc       TEXT
,{{toLower child}}_created    INTEGER
,{{toLower child}}_updated    INTEGER
,UNIQUE({{toLower child}}_{{toLower parent}}_id, {{toLower child}}_slug)
);

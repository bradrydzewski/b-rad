CREATE INDEX IF NOT EXISTS index_{{toLower child}}_{{toLower project}}
ON {{toLower child}}s({{toLower child}}_{{toLower project}}_id);

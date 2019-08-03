CREATE INDEX IF NOT EXISTS index_{{toLower child}}_{{toLower parent}}
ON {{toLower child}}s({{toLower child}}_{{toLower parent}}_id);

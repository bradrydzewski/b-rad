CREATE INDEX IF NOT EXISTS index_{{toLower parent}}_{{toLower project}}
ON {{toLower parent}}s({{toLower parent}}_{{toLower project}}_id);

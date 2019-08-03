CREATE INDEX IF NOT EXISTS index_members_{{toLower project}}
ON members(member_{{toLower project}}_id);

-- name: create-table-members

CREATE TABLE IF NOT EXISTS members (
 member_{{toLower project}}_id INTEGER
,member_user_id    INTEGER
,member_role       INTEGER
,PRIMARY KEY(member_{{toLower project}}_id, member_user_id)
);

-- name: create-index-members-{{toLower project}}-id

CREATE INDEX IF NOT EXISTS index_members_{{toLower project}} ON members(member_{{toLower project}}_id)

-- name: create-index-members-user-id

CREATE INDEX IF NOT EXISTS index_members_user ON members(member_user_id)

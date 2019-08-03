CREATE TABLE IF NOT EXISTS members (
 member_{{toLower project}}_id INTEGER
,member_user_id    INTEGER
,member_role       INTEGER
,PRIMARY KEY(member_{{toLower project}}_id, member_user_id)
);


-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE "Group2" (
  id SERIAL PRIMARY KEY,
  parent_id integer REFERENCES "Group2"(id),
  name       varchar(200),
  organization_id integer REFERENCES organization(id),
  created_at timestamp with time zone not null,
  updated_at timestamp with time zone not null,
  tmp_club_id integer
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS "Group2";

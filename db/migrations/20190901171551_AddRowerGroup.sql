
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE rower_group (
     id serial primary key,
     rower_id integer references rower(id),
     group_id integer references "Group"(id),
     created_at TIMESTAMP WITHOUT TIME ZONE,
     updated_at TIMESTAMP WITHOUT TIME ZONE
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE rower_group;
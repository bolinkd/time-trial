
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE shell ADD COLUMN group_id integer REFERENCES "Group2"(id);
UPDATE shell SET group_id = (SELECT g2.id FROM "Group2" g2 WHERE g2.tmp_club_id = club_id);
ALTER TABLE shell DROP COLUMN club_id;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE shell ADD COLUMN club_id INTEGER REFERENCES club(id);
UPDATE shell SET club_id = (SELECT g2.tmp_club_id FROM "Group2" g2 WHERE g2.id = shell.group_id);
ALTER TABLE shell DROP COLUMN group_id;
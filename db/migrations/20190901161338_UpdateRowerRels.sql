
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE rower DROP COLUMN group_id;
ALTER TABLE rower ADD COLUMN organization_id INTEGER REFERENCES organization(id);
UPDATE rower SET organization_id = 1;

DROP TABLE "Group";
DROP TABLE club;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
CREATE TABLE IF NOT EXISTS Club (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(200),
    "abbreviation" VARCHAR(10),
    "organization_id" INTEGER REFERENCES Organization(id),
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS "Group" (
   "id" SERIAL PRIMARY KEY,
   "name" VARCHAR(200),
   "club_id" INTEGER REFERENCES Club(id),
   "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
   "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL
);

ALTER TABLE rower ADD COLUMN group_id INTEGER REFERENCES "Group"(id);
UPDATE rower r SET group_id =
    ( SELECT min(id) FROM "Group2" g2 WHERE g2.organization_id = r.organization_id);
ALTER TABLE rower DROP COLUMN organization_id;
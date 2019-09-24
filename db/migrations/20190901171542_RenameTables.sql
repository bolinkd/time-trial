
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE "Group2" RENAME TO "Group";

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE "Group" RENAME TO "Group2";

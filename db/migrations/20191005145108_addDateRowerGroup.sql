
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE rower_group
    ADD COLUMN IF NOT EXISTS start_date DATE DEFAULT now(),
    ADD COLUMN IF NOT EXISTS end_date DATE DEFAULT to_date((date_part('year', current_date)+1 || '-08-31'), 'YYYY-MM-DD');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE rower_group
    DROP COLUMN IF EXISTS start_date,
    DROP COLUMN IF EXISTS end_date;

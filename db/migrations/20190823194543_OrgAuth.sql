
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS Organization_Auth (
    "id" SERIAL PRIMARY KEY,
    "organization_id" INTEGER REFERENCES organization(id),
    "token" VARCHAR(50),
    "phrase" VARCHAR(50),
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS OrganizationAuth;
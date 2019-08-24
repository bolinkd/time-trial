
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS Organization (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(200),
  "abbreviation" VARCHAR(10),
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL
);

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

CREATE TABLE IF NOT EXISTS Rower (
  "id" SERIAL PRIMARY KEY,
  "first_name" VARCHAR(50),
  "last_name" VARCHAR(50),
  "group_id" INTEGER REFERENCES "Group"(id),
  "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS Shell (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(200),
  "type" INTEGER,
  "club_id" INTEGER REFERENCES Club(id),
  "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS Rental (
  "id" SERIAL PRIMARY KEY,
  "shell_id" INTEGER REFERENCES Shell(id),
  "out_time" TIMESTAMP WITHOUT TIME ZONE,
  "in_time" TIMESTAMP WITHOUT TIME ZONE,
  "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS Rental_Rowers (
 "id" SERIAL PRIMARY KEY,
 "rental_id" INTEGER REFERENCES Rental(id),
 "rower_id" INTEGER REFERENCES Rower(id),
 "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
 "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS Rental_Rowers;
DROP TABLE IF EXISTS Rental;
DROP TABLE IF EXISTS Shell;
DROP TABLE IF EXISTS Rower;
DROP TABLE IF EXISTS "Group";
DROP TABLE IF EXISTS Club;
DROP TABLE IF EXISTS Organization;


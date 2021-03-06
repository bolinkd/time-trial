
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS Time_Trial (
  "id" SERIAL PRIMARY KEY,
  "date" DATE NOT NULL,
  "start_time" TIMESTAMP,
  "end_time" TIMESTAMP,
  "timing_status" INTEGER,
  "timers" INTEGER,
  "distance" DOUBLE PRECISION,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS Boat (
  "id"            SERIAL PRIMARY KEY,
  "time_trial_id" INTEGER REFERENCES Time_Trial(id),
  "bow_marker"    INTEGER,
  "name"          VARCHAR(50) NOT NULL,
  "start"         INTEGER,
  "end"           INTEGER,
  "time"          INTEGER,
  "created_at"    TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at"    TIMESTAMP WITH TIME ZONE NOT NULL,
  UNIQUE(time_trial_id, bow_marker)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS Boat;
DROP TABLE IF EXISTS Time_Trial;

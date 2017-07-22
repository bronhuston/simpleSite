-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE users
  ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE users
  DROP COLUMN created_at;

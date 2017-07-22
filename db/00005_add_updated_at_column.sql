-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE users
  ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
ON UPDATE CURRENT_TIMESTAMP;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE users
  DROP COLUMN updated_at;

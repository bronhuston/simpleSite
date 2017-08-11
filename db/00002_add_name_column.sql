-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE users ADD COLUMN name VARCHAR(100);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE users DROP COLUMN name;

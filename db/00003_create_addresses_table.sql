-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE addresses (
  id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
  addr_line_1 VARCHAR(100),
  addr_line_2 VARCHAR(100),
  city        VARCHAR(100),
  state       VARCHAR(2),
  zip_5       INTEGER(5),
  user_id     INT UNSIGNED NOT NULL,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE addresses;

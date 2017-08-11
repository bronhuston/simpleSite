-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users
(
  id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (id),
  username    VARCHAR(50)  NOT NULL,
  age         INT,
  description VARCHAR(250)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;

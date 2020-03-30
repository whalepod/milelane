-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE tasks ADD COLUMN expires_at DATETIME AFTER completed_at;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE tasks DROP COLUMN expires_at;


-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER INDEX ADD COLUMN access_token VARCHAR(40)

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE DROP COLUMN access_token

ALTER TABLE user ADD COLUMN access_token VARCHAR(40)
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user DROP COLUMN access_token

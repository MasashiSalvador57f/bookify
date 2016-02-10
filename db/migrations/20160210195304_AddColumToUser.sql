
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `user` ADD COLUMN email VARCHAR(255) NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `user` DROP COLUMN email;
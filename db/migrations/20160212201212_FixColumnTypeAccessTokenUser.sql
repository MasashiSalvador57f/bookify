
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `user` CHANGE COLUMN access_token access_token VARCHAR(64);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `user` CHANGE COLUMN access_token access_token VARCHAR(40);

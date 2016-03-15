
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `user` ADD UNIQUE index_on_access_token (`access_token`)

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `user` DROP INDEX index_on_access_token (`access_token`)

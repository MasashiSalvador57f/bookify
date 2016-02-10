
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `user` MODIFY COLUMN facebook_id VARCHAR(100) NOT NULL;
ALTER TABLE `user` ADD UNIQUE index_on_facebok_id (`facebook_id`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `user` DROP INDEX index_on_facebook_id;
ALTER TABLE `user` MODIFY COLUMN facebook_id TEXT NOT NULL;

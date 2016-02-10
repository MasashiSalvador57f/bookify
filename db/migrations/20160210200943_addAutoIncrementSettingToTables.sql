
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `user`    MODIFY COLUMN id INT UNSIGNED NOT NULL AUTO_INCREMENT;
ALTER TABLE `book`    MODIFY COLUMN id INT UNSIGNED NOT NULL AUTO_INCREMENT;
ALTER TABLE `article` MODIFY COLUMN id INT UNSIGNED NOT NULL AUTO_INCREMENT;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `user`    MODIFY COLUMN id INT UNSIGNED NOT NULL;
ALTER TABLE `book`    MODIFY COLUMN id INT UNSIGNED NOT NULL;
ALTER TABLE `article` MODIFY COLUMN id INT UNSIGNED NOT NULL;

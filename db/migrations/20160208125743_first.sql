
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `book` (
  `id`            INT(10) UNSIGNED NOT NULL,
  `name`          VARCHAR(100)     NOT NULL DEFAULT '',
  `user_id`       INT(10) UNSIGNED NOT NULL,
  `created_at`    DATETIME         NOT NULL,
  `updated_at`    DATETIME         NOT NULL,
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user` (
  `id`                 INT(10) UNSIGNED NOT NULL,
  `name`               VARCHAR(30)      NOT NULL,
  `facebook_id`        TEXT             NOT NULL,
  `created_at`         DATETIME         NOT NULL,
  `updated_at`         DATETIME         NOT NULL,
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `article` (
  `id`                 INT(10) UNSIGNED NOT NULL,
  `title`              VARCHAR(255)     NOT NULL,
  `source_url`         VARCHAR(255)     NOT NULL,
  `created_at`         DATETIME         NOT NULL,
  `updated_at`         DATETIME         NOT NULL,
  PRIMARY KEY(`id`),
  UNIQUE KEY(`source_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `book_article` (
  `book_id`            INT(10) UNSIGNED NOT NULL,
  `article_id`         INT(10) UNSIGNED NOT NULL,
  `created_at`         DATETIME         NOT NULL,
  `updated_at`         DATETIME         NOT NULL,
  PRIMARY KEY(`book_id`, `article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `user_book_article` (
  `user_id`            INT(10)    UNSIGNED NOT NULL,
  `book_id`            INT(10)    UNSIGNED NOT NULL,
  `article_id`         INT(10)    UNSIGNED NOT NULL,
  `status`             TINYINT(3) UNSIGNED NOT NULL,
  PRIMARY KEY(`user_id`, `book_id`, `article_id`),
  KEY(`book_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE `book`;
DROP TABLE `user`;
DROP TABLE `article`;
DROP TABLE `book_article`;
DROP TABLE `user_book_article`;

-- SQL section 'Down' is executed when this migration is rolled back


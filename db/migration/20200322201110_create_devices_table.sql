-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `devices` (
  `uuid`         VARCHAR(36)  NOT NULL PRIMARY KEY,
  `device_token` VARCHAR(255) NOT NULL,
  `type`         INT UNSIGNED NOT NULL,
  `created_at`   DATETIME     NOT NULL,
  `updated_at`   DATETIME     NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `devices` ADD INDEX `idx_device_device_token`(`device_token`);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `devices`;

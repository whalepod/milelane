-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `device_tasks` (
  `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `device_uuid` VARCHAR(36)  NOT NULL,
  `task_id`     INT UNSIGNED NOT NULL,
  `created_at`  DATETIME     NOT NULL,
  `updated_at`  DATETIME     NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `device_tasks` ADD INDEX `idx_device_task_device_uuid`(`device_uuid`);
ALTER TABLE `device_tasks` ADD INDEX `idx_device_task_task_id`(`task_id`);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `device_tasks`;

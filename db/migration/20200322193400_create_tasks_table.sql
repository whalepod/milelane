-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `tasks` (
  `id`           INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title`        VARCHAR(255) NOT NULL,
  `type`         INT UNSIGNED NOT NULL DEFAULT 0,
  `completed_at` DATETIME,
  `created_at`   DATETIME     NOT NULL,
  `updated_at`   DATETIME     NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `tasks` ADD INDEX `idx_task_title`(`title`);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `tasks`;

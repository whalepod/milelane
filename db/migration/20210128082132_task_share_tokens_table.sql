-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `task_share_tokens` (
  `token`              INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `task_id`            VARCHAR(255) NOT NULL,
  `permission_type`    VARCHAR(36) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `task_share_tokens`;

USE `milelane`;

SET CHARACTER_SET_CLIENT = utf8mb4;
SET CHARACTER_SET_CONNECTION = utf8mb4;

CREATE TABLE IF NOT EXISTS `notes` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` VARCHAR(255),
  `body` TEXT NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `devices` (
  `uuid` VARCHAR(36) NOT NULL PRIMARY KEY,
  `device_token` VARCHAR(255) NOT NULL,
  `type` INT UNSIGNED NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `tasks` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` VARCHAR(255) NOT NULL,
  `type` INT UNSIGNED NOT NULL DEFAULT 0,
  `completed_at` DATETIME,
  `deadlined_at` DATETIME,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `tasks` ADD INDEX `idx_task_title`(`title`);

CREATE TABLE IF NOT EXISTS `task_relations` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `ancestor_id` INT UNSIGNED NOT NULL,
  `descendant_id` INT UNSIGNED NOT NULL,
  `path_length` INT UNSIGNED NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `task_relations` ADD INDEX `idx_task_relation_ancestor_id`(`ancestor_id`);
ALTER TABLE `task_relations` ADD INDEX `idx_task_relation_descendant_id`(`descendant_id`);

-- TODO：名前考えたい
CREATE TABLE IF NOT EXISTS `task_action_schedules` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` VARCHAR(255) NOT NULL,
  `task_id` INT UNSIGNED NOT NULL,
  `starts_at` DATETIME,
  `expires_at` DATETIME,
  `created_at` DATETIME NOT NULL,
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `tasks` ADD INDEX `idx_task_title`(`title`);

CREATE TABLE IF NOT EXISTS `device_tasks` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `device_uuid` VARCHAR(36) NOT NULL,
  `task_id` INT UNSIGNED NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `device_tasks` ADD INDEX `idx_device_task_device_uuid`(`device_uuid`);
ALTER TABLE `device_tasks` ADD INDEX `idx_device_task_task_id`(`task_id`);

-- githubのようにprojectsやrepositoriesを作って、それにぶら下げて管理する
-- 状況によってタスクの採番が変わる前提で扱う（devicesから見たtaskのidと、認証後のusersから見たtaskのidが異なりうる)

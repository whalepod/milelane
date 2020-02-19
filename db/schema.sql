CREATE TABLE IF NOT EXISTS `tasks` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` VARCHAR(255) NOT NULL,
  `type` INT UNSIGNED NOT NULL DEFAULT 0,
  `completed_at` DATETIME,
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

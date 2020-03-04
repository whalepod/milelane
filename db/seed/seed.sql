-- device
INSERT INTO `devices` (`id`, `device_id`, `type`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 'dc625158-a9e9-4b7c-b15a-89991b013147', 0, NOW(), NOW());

-- task uncompleted.
INSERT INTO `tasks` (`title`, `type`, `created_at`, `updated_at`) VALUES ('テスト', 0, NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (1, 1, 1, NOW(), NOW());

-- task completed.
INSERT INTO `tasks` (`title`, `type`, `completed_at`, `created_at`, `updated_at`) VALUES ('テスト', 0, NOW(), NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (2, 2, 1, NOW(), NOW());

-- task simple tree ( trunk - branch - leaf ).
-- task which is trunk.
INSERT INTO `tasks` (`title`, `type`, `created_at`, `updated_at`) VALUES ('trunk', 0, NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (3, 3, 1, NOW(), NOW());
-- task which is branch.
INSERT INTO `tasks` (`title`, `type`, `created_at`, `updated_at`) VALUES ('branch', 0, NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (4, 4, 1, NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (3, 4, 2, NOW(), NOW());
-- task which is leaf.
INSERT INTO `tasks` (`title`, `type`, `created_at`, `updated_at`) VALUES ('leaf', 0, NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (5, 5, 1, NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (4, 5, 2, NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (3, 5, 3, NOW(), NOW());
-- task which is branch-2.
INSERT INTO `tasks` (`title`, `type`, `created_at`, `updated_at`) VALUES ('branch-2', 0, NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (6, 6, 1, NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (3, 6, 2, NOW(), NOW());

-- connect tasks with device.
INSERT INTO `device_tasks` (`device_id`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 1, NOW(), NOW());
INSERT INTO `device_tasks` (`device_id`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 2, NOW(), NOW());
INSERT INTO `device_tasks` (`device_id`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 3, NOW(), NOW());
INSERT INTO `device_tasks` (`device_id`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 4, NOW(), NOW());
INSERT INTO `device_tasks` (`device_id`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 5, NOW(), NOW());
INSERT INTO `device_tasks` (`device_id`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 6, NOW(), NOW());

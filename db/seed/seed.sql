-- device
INSERT INTO `devices` (`uuid`, `device_token`, `type`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 'dc625158-a9e9-4b7c-b15a-89991b013147', 0, NOW(), NOW());

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

-- task which is lane.
INSERT INTO `tasks` (`title`, `type`, `created_at`, `updated_at`) VALUES ('テスト', 10, NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (7, 7, 1, NOW(), NOW());

-- task has expires_at ( expires in about 3 month ).
INSERT INTO `tasks` (`title`, `type`, `expires_at`, `created_at`, `updated_at`) VALUES ('テスト', 0, DATE_ADD(NOW(), INTERVAL 90 DAY), NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (8, 8, 1, NOW(), NOW());

-- task already expired.
INSERT INTO `tasks` (`title`, `type`, `expires_at`, `created_at`, `updated_at`) VALUES ('テスト', 0, DATE_SUB(NOW(), INTERVAL 10 DAY), NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (9, 9, 1, NOW(), NOW());

-- task has starts_at ( starts in 10 days ).
INSERT INTO `tasks` (`title`, `type`, `starts_at`, `created_at`, `updated_at`) VALUES ('テスト', 0, DATE_ADD(NOW(), INTERVAL 10 DAY), NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (10, 10, 1, NOW(), NOW());

-- task already started.
INSERT INTO `tasks` (`title`, `type`, `starts_at`, `created_at`, `updated_at`) VALUES ('テスト', 0, DATE_SUB(NOW(), INTERVAL 10 DAY), NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (11, 11, 1, NOW(), NOW());

-- task has starts_at and expires_at ( starts after 10 days, expires after 30 days. ).
INSERT INTO `tasks` (`title`, `type`, `starts_at`, `expires_at`, `created_at`, `updated_at`) VALUES ('テスト', 0, DATE_ADD(NOW(), INTERVAL 10 DAY), DATE_ADD(NOW(), INTERVAL 30 DAY), NOW(), NOW());
INSERT INTO `task_relations` (`ancestor_id`, `descendant_id`, `path_length`, `created_at`, `updated_at`) VALUES (12, 12, 1, NOW(), NOW());

-- connect tasks with device.
INSERT INTO `device_tasks` (`device_uuid`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 1, NOW(), NOW());
INSERT INTO `device_tasks` (`device_uuid`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 2, NOW(), NOW());
INSERT INTO `device_tasks` (`device_uuid`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 3, NOW(), NOW());
INSERT INTO `device_tasks` (`device_uuid`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 4, NOW(), NOW());
INSERT INTO `device_tasks` (`device_uuid`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 5, NOW(), NOW());
INSERT INTO `device_tasks` (`device_uuid`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 6, NOW(), NOW());
INSERT INTO `device_tasks` (`device_uuid`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 7, NOW(), NOW());
INSERT INTO `device_tasks` (`device_uuid`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 8, NOW(), NOW());
INSERT INTO `device_tasks` (`device_uuid`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 9, NOW(), NOW());
INSERT INTO `device_tasks` (`device_uuid`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 10, NOW(), NOW());
INSERT INTO `device_tasks` (`device_uuid`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 11, NOW(), NOW());
INSERT INTO `device_tasks` (`device_uuid`, `task_id`, `created_at`, `updated_at`) VALUES ('60982a48-9328-441f-805b-d3ab0cad9e1f', 12, NOW(), NOW());

-- connect tasks with device.
INSERT INTO `task_share_tokens` (`token`, `task_id`, `permission_type`) VALUES ('abcdefghijklmn', 1, `read`);
USE `milelane`;

-- Avoid 文字化け of Japanese words when inserting data.
SET CHARACTER_SET_CLIENT = utf8mb4;
SET CHARACTER_SET_CONNECTION = utf8mb4;

INSERT INTO notes (
  title,
  body,
  created_at,
  updated_at
) VALUES (
  'タイトルありのノート',
  'タイトルありのノートの本文',
  NOW(),
  NOW()
);

INSERT INTO notes (
  body,
  created_at,
  updated_at
) VALUES (
  'タイトルなしのノートの本文',
  NOW(),
  NOW()
);

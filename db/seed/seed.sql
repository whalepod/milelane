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

INSERT INTO devices (
  uuid,
  device_token,
  type,
  created_at,
  updated_at
)
VALUES (
  '60982a48-9328-441f-805b-d3ab0cad9e1f',
  'dc625158-a9e9-4b7c-b15a-89991b013147',
  0,
  NOW(),
  NOW()
);

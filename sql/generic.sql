-- name: GetAll
SELECT * FROM %{table} OFFSET $1 LIMIT $2;

-- name: GetIDs
SELECT id FROM %{table} OFFSET $1 LIMIT $2;

-- name: TotalNumberOfEntities
SELECT COUNT(*) FROM %{table};

-- name: GetByID
SELECT %{columns} FROM %{table} WHERE id = $1;

-- name: Delete
DELETE FROM %{table} WHERE id = $1;

-- name: Insert
INSERT INTO %{table} (%{columns}) VALUES (%{values});

-- name: Update
UPDATE %{table} %{sets} WHERE id = $1;

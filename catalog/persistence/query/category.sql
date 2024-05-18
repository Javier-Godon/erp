-- name: CreateCategory :one
INSERT INTO catalog.category(category_id, category_name, category_description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateCategory :one
UPDATE catalog.category
SET category_name=$2,
    category_description=$3
WHERE category_id = $1
RETURNING *;

-- name: FindCategoryById :one
SELECT category_id, category_name, category_description
FROM catalog.category
WHERE category_id = $1
LIMIT 1;

-- name: FindCategoryByIdForUpdate :one
SELECT category_id, category_name, category_description
FROM catalog.category
WHERE category_id = $1
LIMIT 1 FOR NO KEY UPDATE;

-- name: FindAllCategories :many
SELECT category_id, category_name, category_description
FROM catalog.category
LIMIT $1 OFFSET $2;


-- name: DeleteCategoryById :exec
DELETE
FROM catalog.category
WHERE category_id = $1;
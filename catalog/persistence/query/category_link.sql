-- name: CreateCategoryLink :one
INSERT INTO catalog.category_link (category_link_id, main_category_id, linked_category_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateCategoryLink :one
UPDATE catalog.category_link
SET main_category_id=$2,
    linked_category_id=$3
WHERE category_link_id = $1
RETURNING *;

-- name: FindCategoryLinkById :one
SELECT category_link_id, main_category_id, linked_category_id
FROM catalog.category_link
WHERE category_link_id = $1
LIMIT 1;

-- name: FindCategoryLinkByLinkedCategoryId :one
SELECT category_link_id, main_category_id, linked_category_id
FROM catalog.category_link
WHERE linked_category_id = $1
LIMIT 1;

-- name: FindCategoryLinkByIdForUpdate :one
SELECT category_link_id, main_category_id, linked_category_id
FROM catalog.category_link
WHERE category_link_id = $1
LIMIT 1 FOR NO KEY UPDATE;

-- name: FindCategoryLinkByLinkedCategoryIdForUpdate :one
SELECT category_link_id, main_category_id, linked_category_id
FROM catalog.category_link
WHERE linked_category_id = $1
LIMIT 1 FOR NO KEY UPDATE;


-- name: FindAllCategoryLinks :many
SELECT category_link_id, main_category_id, linked_category_id
FROM catalog.category_link
LIMIT $1 OFFSET $2;

-- name: DeleteCategoryLinkById :exec
DELETE
FROM catalog.category_link
WHERE category_link_id = $1;

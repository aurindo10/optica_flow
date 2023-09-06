-- name: CreateProduct :one
INSERT INTO product (id, name, price, fornecedor, description, brand, created_at, updated_at)
 values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetAllProducts :many
SELECT * FROM product ORDER BY id ASC;

-- name: UpdateProduct :one
UPDATE product 
SET 
name = coalesce(sqlc.narg('name'), name),
price = coalesce(sqlc.narg('price'), price),
fornecedor = coalesce(sqlc.narg('fornecedor'), fornecedor),
description = coalesce(sqlc.narg('description'), description),
brand = coalesce(sqlc.narg('brand'), brand),
updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeletProduct :exec
DELETE FROM product WHERE id = $1;
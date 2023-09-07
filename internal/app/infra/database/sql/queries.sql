-- name: GetAllProducts :many
SELECT * FROM product WHERE company_id = $1 ORDER BY id ASC;

-- name: CreateProduct :one
WITH valid_fornecedor AS (
  SELECT id
  FROM fornecedor
  WHERE id = $1
)
INSERT INTO product (id, name, price, fornecedor_id, description, brand, created_at, updated_at, bar_code, quantity, company_id, who_created_id, who_updated_id)
SELECT $2, $3, $4, $1, $5, $6, current_timestamp, current_timestamp, $7, $8, $9, $10, $11
FROM valid_fornecedor
RETURNING *;


-- name: UpdateProduct :one
WITH valid_fornecedor AS (
  SELECT id
  FROM fornecedor
  WHERE id = $1
)
UPDATE product
SET name = $2,
    price = $3,
    fornecedor_id = $1,
    description = $4,
    brand = $5,
    updated_at = current_timestamp,
    bar_code = $6,
    quantity = $7,
    company_id = $8,
    who_updated_id = $9
FROM valid_fornecedor
WHERE product.id = $10
RETURNING *;

-- name: DeleteProductById :exec
DELETE FROM product
WHERE id = $1;

-- name: GetAllFornecedores :many
SELECT * FROM fornecedor ORDER BY id ASC;

-- name: CreateFornecedor :one
INSERT INTO fornecedor (id, name, telefone, email, adress, company_id, who_created_id, who_updated_id, cnpj)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateFornecedor :one
UPDATE fornecedor
SET name = $1,
    telefone = $2,
    email = $3,
    adress = $4,
    company_id = $5,
    who_updated_id = $6,
    cnpj = $7
WHERE id = $8
RETURNING *;

-- name: DeleteFornecedorById :exec
DELETE FROM fornecedor
WHERE id = $1;

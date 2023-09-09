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
UPDATE product 
SET 
  name = COALESCE(sqlc.narg('name'), name),
  price = COALESCE(sqlc.narg('price'), price),
  fornecedor_id = COALESCE(sqlc.narg('fornecedor_id'), fornecedor_id),
  description = COALESCE(sqlc.narg('description'), description),
  brand = COALESCE(sqlc.narg('brand'), brand),
  updated_at = current_timestamp,
  bar_code = COALESCE(sqlc.narg('bar_code'), bar_code),
  quantity = COALESCE(sqlc.narg('quantity'), quantity),
  who_updated_id = COALESCE(sqlc.narg('who_updated_id'), who_updated_id)
WHERE id = $1
RETURNING id, name, price, fornecedor_id, description, brand, created_at, updated_at, bar_code, quantity, company_id, who_created_id, who_updated_id;



-- name: DeleteProductById :exec
DELETE FROM product
WHERE id = $1;

-- name: FindAllFornecedores :many
SELECT * FROM fornecedor WHERE company_id = $1 ORDER BY id ASC;

-- name: CreateFornecedor :one
INSERT INTO fornecedor (id, name, telefone, email, adress, company_id, who_created_id, who_updated_id, cnpj)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: DeleteFornecedorById :exec
DELETE FROM fornecedor
WHERE id = $1;

-- name: GetProductByID :one
SELECT * FROM product WHERE id = $1 LIMIT 1;

-- name: GetFornecedorByID :one
SELECT * FROM fornecedor WHERE id = $1 LIMIT 1;

-- name: UpdateFornecedor :one
UPDATE fornecedor 
SET 
  name = COALESCE(sqlc.narg('name'), name),
  telefone = COALESCE(sqlc.narg('telefone'), telefone),
  email = COALESCE(sqlc.narg('email'), email),
  adress = COALESCE(sqlc.narg('adress'), adress),
  company_id = COALESCE(sqlc.narg('company_id'), company_id),
  who_updated_id = COALESCE(sqlc.narg('who_updated_id'), who_updated_id),
  cnpj = COALESCE(sqlc.narg('cnpj'), cnpj)
WHERE id = $1
RETURNING id, name, telefone, email, adress, company_id, who_created_id, who_updated_id, cnpj;


-- name: CreateClient :one
INSERT INTO client (id, full_name, telefone, cpf, created_at, updated_at, email, birth_date, adress, gender, city, seller_id, company_id, who_created_id, who_updated_id)
VALUES ($1, $2, $3, $4, current_timestamp, current_timestamp, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING *;


-- name: FindClientsByCompanyid :many
SELECT * FROM client WHERE company_id = $1 ORDER BY id ASC;

-- name: UpdateClientById :one
UPDATE client
SET
  full_name = COALESCE(sqlc.narg('full_name'), full_name),
  telefone = COALESCE(sqlc.narg('telefone'), telefone),
  cpf = COALESCE(sqlc.narg('cpf'), cpf),
  updated_at = current_timestamp,
  email = COALESCE(sqlc.narg('email'), email),
  birth_date = COALESCE(sqlc.narg('birth_date'), birth_date),
  adress = COALESCE(sqlc.narg('adress'), adress),
  gender = COALESCE(sqlc.narg('gender'), gender),
  city = COALESCE(sqlc.narg('city'), city),
  seller_id = COALESCE(sqlc.narg('seller_id'), seller_id),
  company_id = COALESCE(sqlc.narg('company_id'), company_id),
  who_updated_id = COALESCE(sqlc.narg('who_updated_id'), who_updated_id)
WHERE id = $1
RETURNING id, full_name, telefone, cpf, created_at, updated_at, email, birth_date, adress, gender, city, seller_id, company_id, who_created_id, who_updated_id;


-- name: FindOneClientById :one
SELECT * FROM client WHERE id = $1 LIMIT 1;

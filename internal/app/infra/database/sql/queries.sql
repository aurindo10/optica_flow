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


-- name: DeleteOneClient :exec
DELETE FROM client WHERE id = $1;

-- name: CreateOrder :one
WITH valid_client AS (
  SELECT id
  FROM client
  WHERE id = $1
)
INSERT INTO orders (id, product_name, quantity, order_date, who_created_id, who_updated_id, client_id, company_id, fase)
SELECT $2, $3, $4, current_timestamp, $5, $6, $1, $7, $8
FROM valid_client
RETURNING *;

-- name: FindOneOrderById :one
SELECT * FROM orders WHERE id = $1 LIMIT 1;

-- name: UpdateOneOrder :one
UPDATE orders
SET
  product_name = COALESCE(sqlc.narg('product_name'), product_name),
  quantity = COALESCE(sqlc.narg('quantity'), quantity),
  who_updated_id = COALESCE(sqlc.narg('who_updated_id'), who_updated_id),
  fase = COALESCE(sqlc.narg('fase'), fase)
WHERE id = $1
RETURNING id, product_name, quantity, order_date, who_created_id, who_updated_id, client_id, company_id, fase;

-- name: FindAllORdersByCompanyid :many
SELECT * FROM orders WHERE company_id = $1 ORDER BY id ASC;

-- name: DeleteOrderById :exec
DELETE FROM orders WHERE id = $1;

-- name: CreateProductOrder :one
WITH valid_product AS (
  SELECT id FROM product WHERE id = $1
),
valid_order AS (
  SELECT id FROM orders WHERE id = $2
)
INSERT INTO product_order (id, amout, product_id, order_id)
SELECT $3, $4, $1, $2
FROM valid_product, valid_order
WHERE valid_product.id IS NOT NULL AND valid_order.id IS NOT NULL
RETURNING *;


-- name: FindAllProductOrdersByOrderId :many
SELECT * FROM product_order WHERE order_id = $1 ORDER BY id ASC;

-- name: UpdateProductOrder :one
UPDATE product_order
SET
  amout = COALESCE(sqlc.narg('amout'), amout)
WHERE id = $1
RETURNING id, amout, product_id, order_id;

-- name: DeleteProductOrderById :exec
DELETE FROM product_order WHERE id = $1;

-- name: FindProductOrderById :one
SELECT * FROM product_order WHERE id = $1 LIMIT 1;

-- name: CreatePoints :one
INSERT INTO points (id, name, description, active, ammount, created_at,
updated_at, valid_date, company_id, order_id, seller_id)
SELECT $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
WHERE EXISTS (SELECT 1 FROM orders WHERE id = $10)
RETURNING *;

-- name: FindPointsBySellerId :many
SELECT * FROM points WHERE seller_id = $1 ORDER BY id ASC;

-- name: DeletePointsById :exec
DELETE FROM points WHERE id = $1;

-- name: CreateTradePdrouct :one
INSERT INTO trade_product (id, name, description, created_at,
updated_at, company_id, point_ammount, image_url, who_created_id)
SELECT $1, $2, $3, $4, $5, $6, $7, $8, $9
RETURNING *;


-- name: FindAllTradeProducts :many
SELECT * FROM trade_product WHERE company_id = $1 ORDER BY created_at;

-- name: UpdateTradeProduct :one
UPDATE trade_product
SET 
  name = COALESCE(sqlc.narg('name'), name),
  description = COALESCE(sqlc.narg('description'), description),
  updated_at = COALESCE(sqlc.narg('updated_at'), updated_at),
  point_ammount = COALESCE(sqlc.narg('point_ammount'), point_ammount),
  image_url = COALESCE(sqlc.narg('image_url'), image_url)
WHERE id = $1
RETURNING id, name, description, created_at, updated_at, company_id, point_ammount, image_url, who_created_id;

-- name: DeleteTradeProductById :exec
DELETE FROM trade_product WHERE id = $1;


-- name: CreateComission :one
INSERT INTO commission (id, name, description, created_at, updated_at, company_id, who_created_id, who_updated_id, order_id, value )
SELECT $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
WHERE EXISTS (SELECT 1 FROM orders WHERE id = $9)
RETURNING *;

-- name: FindComissionByUserId :many
SELECT * FROM commission WHERE who_created_id = $1 ORDER BY created_at ASC;

-- name: DeleteComissionById :exec
DELETE FROM commission WHERE id = $1;

-- name: UpdateComission :one
UPDATE commission
SET 
  name = COALESCE(sqlc.narg('name'), name),
  description = COALESCE(sqlc.narg('description'), description),
  updated_at = COALESCE(sqlc.narg('updated_at'),updated_at),
  who_updated_id = COALESCE(sqlc.narg('who_updated_id'),who_updated_id),
  value =  COALESCE(sqlc.narg('value'),value)
WHERE id = $1
RETURNING id, name, description, created_at, updated_at, who_created_id, who_updated_id, order_id, value, company_id;


-- name: CreateComissionValue :one
INSERT INTO comission_values (id, percentage, company_id, who_created_id, who_updated_id, type)
SELECT $1, $2, $3, $4, $5, $6
RETURNING *;


-- name: FindAllComissionValue :many
SELECT * FROM comission_values WHERE company_id = $1 ORDER BY type ASC;

-- name: DeleteComissionValueById :exec
DELETE FROM comission_values WHERE id = $1;

-- name: UpdateComissionValue :one
UPDATE comission_values
SET 
  percentage = COALESCE(sqlc.narg('percentage'), percentage),
  who_updated_id = COALESCE(sqlc.narg('who_updated_id'), who_updated_id),
  type = COALESCE(sqlc.narg('type'), type)
WHERE id = $1
RETURNING id, percentage, company_id, who_created_id, who_updated_id, type;

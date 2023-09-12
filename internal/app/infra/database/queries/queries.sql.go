// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: queries.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createClient = `-- name: CreateClient :one
INSERT INTO client (id, full_name, telefone, cpf, created_at, updated_at, email, birth_date, adress, gender, city, seller_id, company_id, who_created_id, who_updated_id)
VALUES ($1, $2, $3, $4, current_timestamp, current_timestamp, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING id, full_name, telefone, cpf, created_at, updated_at, email, birth_date, adress, gender, city, seller_id, company_id, who_created_id, who_updated_id
`

type CreateClientParams struct {
	ID           uuid.UUID `json:"id"`
	FullName     string    `json:"full_name"`
	Telefone     string    `json:"telefone"`
	Cpf          *string   `json:"cpf"`
	Email        string    `json:"email"`
	BirthDate    time.Time `json:"birth_date"`
	Adress       *string   `json:"adress"`
	Gender       *string   `json:"gender"`
	City         *string   `json:"city"`
	SellerID     *string   `json:"seller_id"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) (Client, error) {
	row := q.queryRow(ctx, q.createClientStmt, createClient,
		arg.ID,
		arg.FullName,
		arg.Telefone,
		arg.Cpf,
		arg.Email,
		arg.BirthDate,
		arg.Adress,
		arg.Gender,
		arg.City,
		arg.SellerID,
		arg.CompanyID,
		arg.WhoCreatedID,
		arg.WhoUpdatedID,
	)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Telefone,
		&i.Cpf,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.BirthDate,
		&i.Adress,
		&i.Gender,
		&i.City,
		&i.SellerID,
		&i.CompanyID,
		&i.WhoCreatedID,
		&i.WhoUpdatedID,
	)
	return i, err
}

const createFornecedor = `-- name: CreateFornecedor :one
INSERT INTO fornecedor (id, name, telefone, email, adress, company_id, who_created_id, who_updated_id, cnpj)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, name, telefone, email, adress, company_id, who_created_id, who_updated_id, cnpj
`

type CreateFornecedorParams struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Telefone     string    `json:"telefone"`
	Email        string    `json:"email"`
	Adress       string    `json:"adress"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
	Cnpj         string    `json:"cnpj"`
}

func (q *Queries) CreateFornecedor(ctx context.Context, arg CreateFornecedorParams) (Fornecedor, error) {
	row := q.queryRow(ctx, q.createFornecedorStmt, createFornecedor,
		arg.ID,
		arg.Name,
		arg.Telefone,
		arg.Email,
		arg.Adress,
		arg.CompanyID,
		arg.WhoCreatedID,
		arg.WhoUpdatedID,
		arg.Cnpj,
	)
	var i Fornecedor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Telefone,
		&i.Email,
		&i.Adress,
		&i.CompanyID,
		&i.WhoCreatedID,
		&i.WhoUpdatedID,
		&i.Cnpj,
	)
	return i, err
}

const createOrder = `-- name: CreateOrder :one
WITH valid_client AS (
  SELECT id
  FROM client
  WHERE id = $1
)
INSERT INTO orders (id, product_name, quantity, order_date, who_created_id, who_updated_id, client_id, company_id, fase)
SELECT $2, $3, $4, current_timestamp, $5, $6, $1, $7, $8
FROM valid_client
RETURNING id, product_name, quantity, order_date, who_created_id, who_updated_id, client_id, company_id, fase
`

type CreateOrderParams struct {
	ClientID     *string   `json:"client_id"`
	ID           uuid.UUID `json:"id"`
	ProductName  string    `json:"product_name"`
	Quantity     int32     `json:"quantity"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
	CompanyID    string    `json:"company_id"`
	Fase         string    `json:"fase"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Orders, error) {
	row := q.queryRow(ctx, q.createOrderStmt, createOrder,
		arg.ClientID,
		arg.ID,
		arg.ProductName,
		arg.Quantity,
		arg.WhoCreatedID,
		arg.WhoUpdatedID,
		arg.CompanyID,
		arg.Fase,
	)
	var i Orders
	err := row.Scan(
		&i.ID,
		&i.ProductName,
		&i.Quantity,
		&i.OrderDate,
		&i.WhoCreatedID,
		&i.WhoUpdatedID,
		&i.ClientID,
		&i.CompanyID,
		&i.Fase,
	)
	return i, err
}

const createPoints = `-- name: CreatePoints :one
INSERT INTO points (id, name, description, active, ammount, created_at,
updated_at, valid_date, company_id, order_id, seller_id)
SELECT $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
WHERE EXISTS (SELECT 1 FROM orders WHERE id = $10)
RETURNING id, name, description, active, ammount, created_at, updated_at, valid_date, company_id, seller_id, order_id
`

type CreatePointsParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	Ammount     float64   `json:"ammount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ValidDate   time.Time `json:"valid_date"`
	CompanyID   string    `json:"company_id"`
	OrderID     *string   `json:"order_id"`
	SellerID    string    `json:"seller_id"`
}

func (q *Queries) CreatePoints(ctx context.Context, arg CreatePointsParams) (Points, error) {
	row := q.queryRow(ctx, q.createPointsStmt, createPoints,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Active,
		arg.Ammount,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ValidDate,
		arg.CompanyID,
		arg.OrderID,
		arg.SellerID,
	)
	var i Points
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Active,
		&i.Ammount,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ValidDate,
		&i.CompanyID,
		&i.SellerID,
		&i.OrderID,
	)
	return i, err
}

const createProduct = `-- name: CreateProduct :one
WITH valid_fornecedor AS (
  SELECT id
  FROM fornecedor
  WHERE id = $1
)
INSERT INTO product (id, name, price, fornecedor_id, description, brand, created_at, updated_at, bar_code, quantity, company_id, who_created_id, who_updated_id)
SELECT $2, $3, $4, $1, $5, $6, current_timestamp, current_timestamp, $7, $8, $9, $10, $11
FROM valid_fornecedor
RETURNING id, name, price, fornecedor_id, description, brand, created_at, updated_at, bar_code, quantity, company_id, who_created_id, who_updated_id
`

type CreateProductParams struct {
	FornecedorID *string   `json:"fornecedor_id"`
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Description  string    `json:"description"`
	Brand        string    `json:"brand"`
	BarCode      string    `json:"bar_code"`
	Quantity     int32     `json:"quantity"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.queryRow(ctx, q.createProductStmt, createProduct,
		arg.FornecedorID,
		arg.ID,
		arg.Name,
		arg.Price,
		arg.Description,
		arg.Brand,
		arg.BarCode,
		arg.Quantity,
		arg.CompanyID,
		arg.WhoCreatedID,
		arg.WhoUpdatedID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.FornecedorID,
		&i.Description,
		&i.Brand,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BarCode,
		&i.Quantity,
		&i.CompanyID,
		&i.WhoCreatedID,
		&i.WhoUpdatedID,
	)
	return i, err
}

const createProductOrder = `-- name: CreateProductOrder :one
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
RETURNING id, amout, product_id, order_id
`

type CreateProductOrderParams struct {
	ProductID *string   `json:"product_id"`
	OrderID   *string   `json:"order_id"`
	ID        uuid.UUID `json:"id"`
	Amout     int32     `json:"amout"`
}

func (q *Queries) CreateProductOrder(ctx context.Context, arg CreateProductOrderParams) (ProductOrder, error) {
	row := q.queryRow(ctx, q.createProductOrderStmt, createProductOrder,
		arg.ProductID,
		arg.OrderID,
		arg.ID,
		arg.Amout,
	)
	var i ProductOrder
	err := row.Scan(
		&i.ID,
		&i.Amout,
		&i.ProductID,
		&i.OrderID,
	)
	return i, err
}

const deleteFornecedorById = `-- name: DeleteFornecedorById :exec
DELETE FROM fornecedor
WHERE id = $1
`

func (q *Queries) DeleteFornecedorById(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteFornecedorByIdStmt, deleteFornecedorById, id)
	return err
}

const deleteOneClient = `-- name: DeleteOneClient :exec
DELETE FROM client WHERE id = $1
`

func (q *Queries) DeleteOneClient(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteOneClientStmt, deleteOneClient, id)
	return err
}

const deleteOrderById = `-- name: DeleteOrderById :exec
DELETE FROM orders WHERE id = $1
`

func (q *Queries) DeleteOrderById(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteOrderByIdStmt, deleteOrderById, id)
	return err
}

const deleteProductById = `-- name: DeleteProductById :exec
DELETE FROM product
WHERE id = $1
`

func (q *Queries) DeleteProductById(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteProductByIdStmt, deleteProductById, id)
	return err
}

const deleteProductOrderById = `-- name: DeleteProductOrderById :exec
DELETE FROM product_order WHERE id = $1
`

func (q *Queries) DeleteProductOrderById(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteProductOrderByIdStmt, deleteProductOrderById, id)
	return err
}

const findAllFornecedores = `-- name: FindAllFornecedores :many
SELECT id, name, telefone, email, adress, company_id, who_created_id, who_updated_id, cnpj FROM fornecedor WHERE company_id = $1 ORDER BY id ASC
`

func (q *Queries) FindAllFornecedores(ctx context.Context, companyID string) ([]Fornecedor, error) {
	rows, err := q.query(ctx, q.findAllFornecedoresStmt, findAllFornecedores, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Fornecedor
	for rows.Next() {
		var i Fornecedor
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Telefone,
			&i.Email,
			&i.Adress,
			&i.CompanyID,
			&i.WhoCreatedID,
			&i.WhoUpdatedID,
			&i.Cnpj,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllORdersByCompanyid = `-- name: FindAllORdersByCompanyid :many
SELECT id, product_name, quantity, order_date, who_created_id, who_updated_id, client_id, company_id, fase FROM orders WHERE company_id = $1 ORDER BY id ASC
`

func (q *Queries) FindAllORdersByCompanyid(ctx context.Context, companyID string) ([]Orders, error) {
	rows, err := q.query(ctx, q.findAllORdersByCompanyidStmt, findAllORdersByCompanyid, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Orders
	for rows.Next() {
		var i Orders
		if err := rows.Scan(
			&i.ID,
			&i.ProductName,
			&i.Quantity,
			&i.OrderDate,
			&i.WhoCreatedID,
			&i.WhoUpdatedID,
			&i.ClientID,
			&i.CompanyID,
			&i.Fase,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllProductOrdersByOrderId = `-- name: FindAllProductOrdersByOrderId :many
SELECT id, amout, product_id, order_id FROM product_order WHERE order_id = $1 ORDER BY id ASC
`

func (q *Queries) FindAllProductOrdersByOrderId(ctx context.Context, orderID *string) ([]ProductOrder, error) {
	rows, err := q.query(ctx, q.findAllProductOrdersByOrderIdStmt, findAllProductOrdersByOrderId, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductOrder
	for rows.Next() {
		var i ProductOrder
		if err := rows.Scan(
			&i.ID,
			&i.Amout,
			&i.ProductID,
			&i.OrderID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findClientsByCompanyid = `-- name: FindClientsByCompanyid :many
SELECT id, full_name, telefone, cpf, created_at, updated_at, email, birth_date, adress, gender, city, seller_id, company_id, who_created_id, who_updated_id FROM client WHERE company_id = $1 ORDER BY id ASC
`

func (q *Queries) FindClientsByCompanyid(ctx context.Context, companyID string) ([]Client, error) {
	rows, err := q.query(ctx, q.findClientsByCompanyidStmt, findClientsByCompanyid, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Client
	for rows.Next() {
		var i Client
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.Telefone,
			&i.Cpf,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Email,
			&i.BirthDate,
			&i.Adress,
			&i.Gender,
			&i.City,
			&i.SellerID,
			&i.CompanyID,
			&i.WhoCreatedID,
			&i.WhoUpdatedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findOneClientById = `-- name: FindOneClientById :one
SELECT id, full_name, telefone, cpf, created_at, updated_at, email, birth_date, adress, gender, city, seller_id, company_id, who_created_id, who_updated_id FROM client WHERE id = $1 LIMIT 1
`

func (q *Queries) FindOneClientById(ctx context.Context, id uuid.UUID) (Client, error) {
	row := q.queryRow(ctx, q.findOneClientByIdStmt, findOneClientById, id)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Telefone,
		&i.Cpf,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.BirthDate,
		&i.Adress,
		&i.Gender,
		&i.City,
		&i.SellerID,
		&i.CompanyID,
		&i.WhoCreatedID,
		&i.WhoUpdatedID,
	)
	return i, err
}

const findOneOrderById = `-- name: FindOneOrderById :one
SELECT id, product_name, quantity, order_date, who_created_id, who_updated_id, client_id, company_id, fase FROM orders WHERE id = $1 LIMIT 1
`

func (q *Queries) FindOneOrderById(ctx context.Context, id uuid.UUID) (Orders, error) {
	row := q.queryRow(ctx, q.findOneOrderByIdStmt, findOneOrderById, id)
	var i Orders
	err := row.Scan(
		&i.ID,
		&i.ProductName,
		&i.Quantity,
		&i.OrderDate,
		&i.WhoCreatedID,
		&i.WhoUpdatedID,
		&i.ClientID,
		&i.CompanyID,
		&i.Fase,
	)
	return i, err
}

const findPointsBySellerId = `-- name: FindPointsBySellerId :many
SELECT id, name, description, active, ammount, created_at, updated_at, valid_date, company_id, seller_id, order_id FROM points WHERE seller_id = $1 ORDER BY id ASC
`

func (q *Queries) FindPointsBySellerId(ctx context.Context, sellerID string) ([]Points, error) {
	rows, err := q.query(ctx, q.findPointsBySellerIdStmt, findPointsBySellerId, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Points
	for rows.Next() {
		var i Points
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Active,
			&i.Ammount,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ValidDate,
			&i.CompanyID,
			&i.SellerID,
			&i.OrderID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findProductOrderById = `-- name: FindProductOrderById :one
SELECT id, amout, product_id, order_id FROM product_order WHERE id = $1 LIMIT 1
`

func (q *Queries) FindProductOrderById(ctx context.Context, id uuid.UUID) (ProductOrder, error) {
	row := q.queryRow(ctx, q.findProductOrderByIdStmt, findProductOrderById, id)
	var i ProductOrder
	err := row.Scan(
		&i.ID,
		&i.Amout,
		&i.ProductID,
		&i.OrderID,
	)
	return i, err
}

const getAllProducts = `-- name: GetAllProducts :many
SELECT id, name, price, fornecedor_id, description, brand, created_at, updated_at, bar_code, quantity, company_id, who_created_id, who_updated_id FROM product WHERE company_id = $1 ORDER BY id ASC
`

func (q *Queries) GetAllProducts(ctx context.Context, companyID string) ([]Product, error) {
	rows, err := q.query(ctx, q.getAllProductsStmt, getAllProducts, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.FornecedorID,
			&i.Description,
			&i.Brand,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.BarCode,
			&i.Quantity,
			&i.CompanyID,
			&i.WhoCreatedID,
			&i.WhoUpdatedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFornecedorByID = `-- name: GetFornecedorByID :one
SELECT id, name, telefone, email, adress, company_id, who_created_id, who_updated_id, cnpj FROM fornecedor WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFornecedorByID(ctx context.Context, id uuid.UUID) (Fornecedor, error) {
	row := q.queryRow(ctx, q.getFornecedorByIDStmt, getFornecedorByID, id)
	var i Fornecedor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Telefone,
		&i.Email,
		&i.Adress,
		&i.CompanyID,
		&i.WhoCreatedID,
		&i.WhoUpdatedID,
		&i.Cnpj,
	)
	return i, err
}

const getProductByID = `-- name: GetProductByID :one
SELECT id, name, price, fornecedor_id, description, brand, created_at, updated_at, bar_code, quantity, company_id, who_created_id, who_updated_id FROM product WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProductByID(ctx context.Context, id uuid.UUID) (Product, error) {
	row := q.queryRow(ctx, q.getProductByIDStmt, getProductByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.FornecedorID,
		&i.Description,
		&i.Brand,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BarCode,
		&i.Quantity,
		&i.CompanyID,
		&i.WhoCreatedID,
		&i.WhoUpdatedID,
	)
	return i, err
}

const updateClientById = `-- name: UpdateClientById :one
UPDATE client
SET
  full_name = COALESCE($2, full_name),
  telefone = COALESCE($3, telefone),
  cpf = COALESCE($4, cpf),
  updated_at = current_timestamp,
  email = COALESCE($5, email),
  birth_date = COALESCE($6, birth_date),
  adress = COALESCE($7, adress),
  gender = COALESCE($8, gender),
  city = COALESCE($9, city),
  seller_id = COALESCE($10, seller_id),
  company_id = COALESCE($11, company_id),
  who_updated_id = COALESCE($12, who_updated_id)
WHERE id = $1
RETURNING id, full_name, telefone, cpf, created_at, updated_at, email, birth_date, adress, gender, city, seller_id, company_id, who_created_id, who_updated_id
`

type UpdateClientByIdParams struct {
	ID           uuid.UUID  `json:"id"`
	FullName     *string    `json:"full_name"`
	Telefone     *string    `json:"telefone"`
	Cpf          *string    `json:"cpf"`
	Email        *string    `json:"email"`
	BirthDate    *time.Time `json:"birth_date"`
	Adress       *string    `json:"adress"`
	Gender       *string    `json:"gender"`
	City         *string    `json:"city"`
	SellerID     *string    `json:"seller_id"`
	CompanyID    *string    `json:"company_id"`
	WhoUpdatedID *string    `json:"who_updated_id"`
}

func (q *Queries) UpdateClientById(ctx context.Context, arg UpdateClientByIdParams) (Client, error) {
	row := q.queryRow(ctx, q.updateClientByIdStmt, updateClientById,
		arg.ID,
		arg.FullName,
		arg.Telefone,
		arg.Cpf,
		arg.Email,
		arg.BirthDate,
		arg.Adress,
		arg.Gender,
		arg.City,
		arg.SellerID,
		arg.CompanyID,
		arg.WhoUpdatedID,
	)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Telefone,
		&i.Cpf,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.BirthDate,
		&i.Adress,
		&i.Gender,
		&i.City,
		&i.SellerID,
		&i.CompanyID,
		&i.WhoCreatedID,
		&i.WhoUpdatedID,
	)
	return i, err
}

const updateFornecedor = `-- name: UpdateFornecedor :one
UPDATE fornecedor 
SET 
  name = COALESCE($2, name),
  telefone = COALESCE($3, telefone),
  email = COALESCE($4, email),
  adress = COALESCE($5, adress),
  company_id = COALESCE($6, company_id),
  who_updated_id = COALESCE($7, who_updated_id),
  cnpj = COALESCE($8, cnpj)
WHERE id = $1
RETURNING id, name, telefone, email, adress, company_id, who_created_id, who_updated_id, cnpj
`

type UpdateFornecedorParams struct {
	ID           uuid.UUID `json:"id"`
	Name         *string   `json:"name"`
	Telefone     *string   `json:"telefone"`
	Email        *string   `json:"email"`
	Adress       *string   `json:"adress"`
	CompanyID    *string   `json:"company_id"`
	WhoUpdatedID *string   `json:"who_updated_id"`
	Cnpj         *string   `json:"cnpj"`
}

func (q *Queries) UpdateFornecedor(ctx context.Context, arg UpdateFornecedorParams) (Fornecedor, error) {
	row := q.queryRow(ctx, q.updateFornecedorStmt, updateFornecedor,
		arg.ID,
		arg.Name,
		arg.Telefone,
		arg.Email,
		arg.Adress,
		arg.CompanyID,
		arg.WhoUpdatedID,
		arg.Cnpj,
	)
	var i Fornecedor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Telefone,
		&i.Email,
		&i.Adress,
		&i.CompanyID,
		&i.WhoCreatedID,
		&i.WhoUpdatedID,
		&i.Cnpj,
	)
	return i, err
}

const updateOneOrder = `-- name: UpdateOneOrder :one
UPDATE orders
SET
  product_name = COALESCE($2, product_name),
  quantity = COALESCE($3, quantity),
  who_updated_id = COALESCE($4, who_updated_id),
  fase = COALESCE($5, fase)
WHERE id = $1
RETURNING id, product_name, quantity, order_date, who_created_id, who_updated_id, client_id, company_id, fase
`

type UpdateOneOrderParams struct {
	ID           uuid.UUID `json:"id"`
	ProductName  *string   `json:"product_name"`
	Quantity     *int32    `json:"quantity"`
	WhoUpdatedID *string   `json:"who_updated_id"`
	Fase         *string   `json:"fase"`
}

func (q *Queries) UpdateOneOrder(ctx context.Context, arg UpdateOneOrderParams) (Orders, error) {
	row := q.queryRow(ctx, q.updateOneOrderStmt, updateOneOrder,
		arg.ID,
		arg.ProductName,
		arg.Quantity,
		arg.WhoUpdatedID,
		arg.Fase,
	)
	var i Orders
	err := row.Scan(
		&i.ID,
		&i.ProductName,
		&i.Quantity,
		&i.OrderDate,
		&i.WhoCreatedID,
		&i.WhoUpdatedID,
		&i.ClientID,
		&i.CompanyID,
		&i.Fase,
	)
	return i, err
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE product 
SET 
  name = COALESCE($2, name),
  price = COALESCE($3, price),
  fornecedor_id = COALESCE($4, fornecedor_id),
  description = COALESCE($5, description),
  brand = COALESCE($6, brand),
  updated_at = current_timestamp,
  bar_code = COALESCE($7, bar_code),
  quantity = COALESCE($8, quantity),
  who_updated_id = COALESCE($9, who_updated_id)
WHERE id = $1
RETURNING id, name, price, fornecedor_id, description, brand, created_at, updated_at, bar_code, quantity, company_id, who_created_id, who_updated_id
`

type UpdateProductParams struct {
	ID           uuid.UUID `json:"id"`
	Name         *string   `json:"name"`
	Price        *float64  `json:"price"`
	FornecedorID *string   `json:"fornecedor_id"`
	Description  *string   `json:"description"`
	Brand        *string   `json:"brand"`
	BarCode      *string   `json:"bar_code"`
	Quantity     *int32    `json:"quantity"`
	WhoUpdatedID *string   `json:"who_updated_id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.queryRow(ctx, q.updateProductStmt, updateProduct,
		arg.ID,
		arg.Name,
		arg.Price,
		arg.FornecedorID,
		arg.Description,
		arg.Brand,
		arg.BarCode,
		arg.Quantity,
		arg.WhoUpdatedID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.FornecedorID,
		&i.Description,
		&i.Brand,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BarCode,
		&i.Quantity,
		&i.CompanyID,
		&i.WhoCreatedID,
		&i.WhoUpdatedID,
	)
	return i, err
}

const updateProductOrder = `-- name: UpdateProductOrder :one
UPDATE product_order
SET
  amout = COALESCE($2, amout)
WHERE id = $1
RETURNING id, amout, product_id, order_id
`

type UpdateProductOrderParams struct {
	ID    uuid.UUID `json:"id"`
	Amout *int32    `json:"amout"`
}

func (q *Queries) UpdateProductOrder(ctx context.Context, arg UpdateProductOrderParams) (ProductOrder, error) {
	row := q.queryRow(ctx, q.updateProductOrderStmt, updateProductOrder, arg.ID, arg.Amout)
	var i ProductOrder
	err := row.Scan(
		&i.ID,
		&i.Amout,
		&i.ProductID,
		&i.OrderID,
	)
	return i, err
}

package repository

import (
	"context"
	"optica_flow/internal/app/domain/product"
	database "optica_flow/internal/app/infra/database/queries"

	"github.com/google/uuid"
)

type ProductRepository struct {
	db *database.Queries
}



func (p *ProductRepository) CreateProduct(product *product.Product) error {
	ctx := context.Background()
	info := database.CreateProductParams{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Brand: 	 	product.Brand,
		FornecedorID: product.FornecedorID,
		BarCode:    product.BarCode,
		Quantity:  product.Quantity,
		CompanyID: product.CompanyID,
		WhoCreatedID: product.WhoCreatedID,
		WhoUpdatedID: product.WhoUpdatedID,
		Description: product.Description,
	}
	if _, err := p.db.CreateProduct(ctx, info); err != nil {
		return err
	}
	return nil
}
func (p *ProductRepository) GetAllProducts(id string) ([]*product.Product, error) {
	ctx := context.Background()
	products, err := p.db.GetAllProducts(ctx, id)
	if err != nil {
		return nil, err
	}
	return p.mapResult(&products), nil
}
func (p *ProductRepository) UpdateProduct(productRequest *product.ProductRequestToUpdate)  (*product.Product, error){
	ctx := context.Background()
	info := database.UpdateProductParams{
		ID : productRequest.ID,
		Name: productRequest.Name,
		Price: productRequest.Price,
		Brand: productRequest.Brand,
		BarCode: productRequest.BarCode,
		Quantity: productRequest.Quantity,
		Description: productRequest.Description,
		FornecedorID: productRequest.FornecedorID,
		WhoUpdatedID: productRequest.WhoUpdatedID,
	}
	updatedProduct, error := p.db.UpdateProduct(ctx, info)
	if error != nil {
		return nil, error
	}
	productUpdated := &product.Product{
		ID: updatedProduct.ID,
		Name: updatedProduct.Name,
		Price: updatedProduct.Price,
		FornecedorID: updatedProduct.FornecedorID,
		Description: updatedProduct.Description,
		Brand: updatedProduct.Brand,
		CreatedAt: updatedProduct.CreatedAt,
		UpdatedAt: updatedProduct.UpdatedAt,
		BarCode: updatedProduct.BarCode,
		Quantity: updatedProduct.Quantity,
		CompanyID: updatedProduct.CompanyID,
		WhoCreatedID: updatedProduct.WhoCreatedID,
		WhoUpdatedID: updatedProduct.WhoUpdatedID,
	}
	 return productUpdated, nil
}
func NewProductRepository(db *database.Queries) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (c *ProductRepository) mapResult(result *[]database.Product) []*product.Product {
	var products []*product.Product
	for _, v := range *result {
		products = append(products, &product.Product{
			ID:          v.ID,
			Name:        v.Name,
			Price:       v.Price,
			Brand: 	 	v.Brand,
			FornecedorID:  v.FornecedorID,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			BarCode:    v.BarCode,
			Quantity:  v.Quantity,
			CompanyID: v.CompanyID,
			WhoCreatedID: v.WhoCreatedID,
			WhoUpdatedID: v.WhoUpdatedID,
		})
	}
	return products
}

func (c *ProductRepository) DeleteProduct(id uuid.UUID) error {
	if error := c.db.DeleteProductById(context.Background(), id); error != nil {
		return error
	}
	return nil
}

func (c *ProductRepository) GetProductById(id uuid.UUID) (*product.Product, error) {
	p, error := c.db.GetProductByID(context.Background(), id)
	if error != nil {
		return nil, error
	}
	response := &product.Product{
		ID:          p.ID,
		Name:        p.Name,
		Price:       p.Price,
		Brand: 	 	p.Brand,
		FornecedorID:  p.FornecedorID,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		BarCode:    p.BarCode,
		Quantity:  p.Quantity,
		CompanyID: p.CompanyID,
		WhoCreatedID: p.WhoCreatedID,
		WhoUpdatedID: p.WhoUpdatedID,
	}
	return response, nil
}
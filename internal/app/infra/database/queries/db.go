// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package database

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createClientStmt, err = db.PrepareContext(ctx, createClient); err != nil {
		return nil, fmt.Errorf("error preparing query CreateClient: %w", err)
	}
	if q.createFornecedorStmt, err = db.PrepareContext(ctx, createFornecedor); err != nil {
		return nil, fmt.Errorf("error preparing query CreateFornecedor: %w", err)
	}
	if q.createOrderStmt, err = db.PrepareContext(ctx, createOrder); err != nil {
		return nil, fmt.Errorf("error preparing query CreateOrder: %w", err)
	}
	if q.createPointsStmt, err = db.PrepareContext(ctx, createPoints); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePoints: %w", err)
	}
	if q.createProductStmt, err = db.PrepareContext(ctx, createProduct); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProduct: %w", err)
	}
	if q.createProductOrderStmt, err = db.PrepareContext(ctx, createProductOrder); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProductOrder: %w", err)
	}
	if q.createTradePdrouctStmt, err = db.PrepareContext(ctx, createTradePdrouct); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTradePdrouct: %w", err)
	}
	if q.deleteFornecedorByIdStmt, err = db.PrepareContext(ctx, deleteFornecedorById); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteFornecedorById: %w", err)
	}
	if q.deleteOneClientStmt, err = db.PrepareContext(ctx, deleteOneClient); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteOneClient: %w", err)
	}
	if q.deleteOrderByIdStmt, err = db.PrepareContext(ctx, deleteOrderById); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteOrderById: %w", err)
	}
	if q.deletePointsByIdStmt, err = db.PrepareContext(ctx, deletePointsById); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePointsById: %w", err)
	}
	if q.deleteProductByIdStmt, err = db.PrepareContext(ctx, deleteProductById); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProductById: %w", err)
	}
	if q.deleteProductOrderByIdStmt, err = db.PrepareContext(ctx, deleteProductOrderById); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProductOrderById: %w", err)
	}
	if q.deleteTradeProductByIdStmt, err = db.PrepareContext(ctx, deleteTradeProductById); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTradeProductById: %w", err)
	}
	if q.findAllFornecedoresStmt, err = db.PrepareContext(ctx, findAllFornecedores); err != nil {
		return nil, fmt.Errorf("error preparing query FindAllFornecedores: %w", err)
	}
	if q.findAllORdersByCompanyidStmt, err = db.PrepareContext(ctx, findAllORdersByCompanyid); err != nil {
		return nil, fmt.Errorf("error preparing query FindAllORdersByCompanyid: %w", err)
	}
	if q.findAllProductOrdersByOrderIdStmt, err = db.PrepareContext(ctx, findAllProductOrdersByOrderId); err != nil {
		return nil, fmt.Errorf("error preparing query FindAllProductOrdersByOrderId: %w", err)
	}
	if q.findAllTradeProductsStmt, err = db.PrepareContext(ctx, findAllTradeProducts); err != nil {
		return nil, fmt.Errorf("error preparing query FindAllTradeProducts: %w", err)
	}
	if q.findClientsByCompanyidStmt, err = db.PrepareContext(ctx, findClientsByCompanyid); err != nil {
		return nil, fmt.Errorf("error preparing query FindClientsByCompanyid: %w", err)
	}
	if q.findOneClientByIdStmt, err = db.PrepareContext(ctx, findOneClientById); err != nil {
		return nil, fmt.Errorf("error preparing query FindOneClientById: %w", err)
	}
	if q.findOneOrderByIdStmt, err = db.PrepareContext(ctx, findOneOrderById); err != nil {
		return nil, fmt.Errorf("error preparing query FindOneOrderById: %w", err)
	}
	if q.findPointsBySellerIdStmt, err = db.PrepareContext(ctx, findPointsBySellerId); err != nil {
		return nil, fmt.Errorf("error preparing query FindPointsBySellerId: %w", err)
	}
	if q.findProductOrderByIdStmt, err = db.PrepareContext(ctx, findProductOrderById); err != nil {
		return nil, fmt.Errorf("error preparing query FindProductOrderById: %w", err)
	}
	if q.getAllProductsStmt, err = db.PrepareContext(ctx, getAllProducts); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllProducts: %w", err)
	}
	if q.getFornecedorByIDStmt, err = db.PrepareContext(ctx, getFornecedorByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetFornecedorByID: %w", err)
	}
	if q.getProductByIDStmt, err = db.PrepareContext(ctx, getProductByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductByID: %w", err)
	}
	if q.updateClientByIdStmt, err = db.PrepareContext(ctx, updateClientById); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateClientById: %w", err)
	}
	if q.updateFornecedorStmt, err = db.PrepareContext(ctx, updateFornecedor); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateFornecedor: %w", err)
	}
	if q.updateOneOrderStmt, err = db.PrepareContext(ctx, updateOneOrder); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateOneOrder: %w", err)
	}
	if q.updateProductStmt, err = db.PrepareContext(ctx, updateProduct); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProduct: %w", err)
	}
	if q.updateProductOrderStmt, err = db.PrepareContext(ctx, updateProductOrder); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProductOrder: %w", err)
	}
	if q.updateTradeProductStmt, err = db.PrepareContext(ctx, updateTradeProduct); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTradeProduct: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createClientStmt != nil {
		if cerr := q.createClientStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createClientStmt: %w", cerr)
		}
	}
	if q.createFornecedorStmt != nil {
		if cerr := q.createFornecedorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createFornecedorStmt: %w", cerr)
		}
	}
	if q.createOrderStmt != nil {
		if cerr := q.createOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createOrderStmt: %w", cerr)
		}
	}
	if q.createPointsStmt != nil {
		if cerr := q.createPointsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPointsStmt: %w", cerr)
		}
	}
	if q.createProductStmt != nil {
		if cerr := q.createProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProductStmt: %w", cerr)
		}
	}
	if q.createProductOrderStmt != nil {
		if cerr := q.createProductOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProductOrderStmt: %w", cerr)
		}
	}
	if q.createTradePdrouctStmt != nil {
		if cerr := q.createTradePdrouctStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTradePdrouctStmt: %w", cerr)
		}
	}
	if q.deleteFornecedorByIdStmt != nil {
		if cerr := q.deleteFornecedorByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteFornecedorByIdStmt: %w", cerr)
		}
	}
	if q.deleteOneClientStmt != nil {
		if cerr := q.deleteOneClientStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteOneClientStmt: %w", cerr)
		}
	}
	if q.deleteOrderByIdStmt != nil {
		if cerr := q.deleteOrderByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteOrderByIdStmt: %w", cerr)
		}
	}
	if q.deletePointsByIdStmt != nil {
		if cerr := q.deletePointsByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePointsByIdStmt: %w", cerr)
		}
	}
	if q.deleteProductByIdStmt != nil {
		if cerr := q.deleteProductByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProductByIdStmt: %w", cerr)
		}
	}
	if q.deleteProductOrderByIdStmt != nil {
		if cerr := q.deleteProductOrderByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProductOrderByIdStmt: %w", cerr)
		}
	}
	if q.deleteTradeProductByIdStmt != nil {
		if cerr := q.deleteTradeProductByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTradeProductByIdStmt: %w", cerr)
		}
	}
	if q.findAllFornecedoresStmt != nil {
		if cerr := q.findAllFornecedoresStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findAllFornecedoresStmt: %w", cerr)
		}
	}
	if q.findAllORdersByCompanyidStmt != nil {
		if cerr := q.findAllORdersByCompanyidStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findAllORdersByCompanyidStmt: %w", cerr)
		}
	}
	if q.findAllProductOrdersByOrderIdStmt != nil {
		if cerr := q.findAllProductOrdersByOrderIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findAllProductOrdersByOrderIdStmt: %w", cerr)
		}
	}
	if q.findAllTradeProductsStmt != nil {
		if cerr := q.findAllTradeProductsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findAllTradeProductsStmt: %w", cerr)
		}
	}
	if q.findClientsByCompanyidStmt != nil {
		if cerr := q.findClientsByCompanyidStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findClientsByCompanyidStmt: %w", cerr)
		}
	}
	if q.findOneClientByIdStmt != nil {
		if cerr := q.findOneClientByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findOneClientByIdStmt: %w", cerr)
		}
	}
	if q.findOneOrderByIdStmt != nil {
		if cerr := q.findOneOrderByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findOneOrderByIdStmt: %w", cerr)
		}
	}
	if q.findPointsBySellerIdStmt != nil {
		if cerr := q.findPointsBySellerIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findPointsBySellerIdStmt: %w", cerr)
		}
	}
	if q.findProductOrderByIdStmt != nil {
		if cerr := q.findProductOrderByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findProductOrderByIdStmt: %w", cerr)
		}
	}
	if q.getAllProductsStmt != nil {
		if cerr := q.getAllProductsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllProductsStmt: %w", cerr)
		}
	}
	if q.getFornecedorByIDStmt != nil {
		if cerr := q.getFornecedorByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFornecedorByIDStmt: %w", cerr)
		}
	}
	if q.getProductByIDStmt != nil {
		if cerr := q.getProductByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductByIDStmt: %w", cerr)
		}
	}
	if q.updateClientByIdStmt != nil {
		if cerr := q.updateClientByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateClientByIdStmt: %w", cerr)
		}
	}
	if q.updateFornecedorStmt != nil {
		if cerr := q.updateFornecedorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateFornecedorStmt: %w", cerr)
		}
	}
	if q.updateOneOrderStmt != nil {
		if cerr := q.updateOneOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateOneOrderStmt: %w", cerr)
		}
	}
	if q.updateProductStmt != nil {
		if cerr := q.updateProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProductStmt: %w", cerr)
		}
	}
	if q.updateProductOrderStmt != nil {
		if cerr := q.updateProductOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProductOrderStmt: %w", cerr)
		}
	}
	if q.updateTradeProductStmt != nil {
		if cerr := q.updateTradeProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTradeProductStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                DBTX
	tx                                *sql.Tx
	createClientStmt                  *sql.Stmt
	createFornecedorStmt              *sql.Stmt
	createOrderStmt                   *sql.Stmt
	createPointsStmt                  *sql.Stmt
	createProductStmt                 *sql.Stmt
	createProductOrderStmt            *sql.Stmt
	createTradePdrouctStmt            *sql.Stmt
	deleteFornecedorByIdStmt          *sql.Stmt
	deleteOneClientStmt               *sql.Stmt
	deleteOrderByIdStmt               *sql.Stmt
	deletePointsByIdStmt              *sql.Stmt
	deleteProductByIdStmt             *sql.Stmt
	deleteProductOrderByIdStmt        *sql.Stmt
	deleteTradeProductByIdStmt        *sql.Stmt
	findAllFornecedoresStmt           *sql.Stmt
	findAllORdersByCompanyidStmt      *sql.Stmt
	findAllProductOrdersByOrderIdStmt *sql.Stmt
	findAllTradeProductsStmt          *sql.Stmt
	findClientsByCompanyidStmt        *sql.Stmt
	findOneClientByIdStmt             *sql.Stmt
	findOneOrderByIdStmt              *sql.Stmt
	findPointsBySellerIdStmt          *sql.Stmt
	findProductOrderByIdStmt          *sql.Stmt
	getAllProductsStmt                *sql.Stmt
	getFornecedorByIDStmt             *sql.Stmt
	getProductByIDStmt                *sql.Stmt
	updateClientByIdStmt              *sql.Stmt
	updateFornecedorStmt              *sql.Stmt
	updateOneOrderStmt                *sql.Stmt
	updateProductStmt                 *sql.Stmt
	updateProductOrderStmt            *sql.Stmt
	updateTradeProductStmt            *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                tx,
		tx:                                tx,
		createClientStmt:                  q.createClientStmt,
		createFornecedorStmt:              q.createFornecedorStmt,
		createOrderStmt:                   q.createOrderStmt,
		createPointsStmt:                  q.createPointsStmt,
		createProductStmt:                 q.createProductStmt,
		createProductOrderStmt:            q.createProductOrderStmt,
		createTradePdrouctStmt:            q.createTradePdrouctStmt,
		deleteFornecedorByIdStmt:          q.deleteFornecedorByIdStmt,
		deleteOneClientStmt:               q.deleteOneClientStmt,
		deleteOrderByIdStmt:               q.deleteOrderByIdStmt,
		deletePointsByIdStmt:              q.deletePointsByIdStmt,
		deleteProductByIdStmt:             q.deleteProductByIdStmt,
		deleteProductOrderByIdStmt:        q.deleteProductOrderByIdStmt,
		deleteTradeProductByIdStmt:        q.deleteTradeProductByIdStmt,
		findAllFornecedoresStmt:           q.findAllFornecedoresStmt,
		findAllORdersByCompanyidStmt:      q.findAllORdersByCompanyidStmt,
		findAllProductOrdersByOrderIdStmt: q.findAllProductOrdersByOrderIdStmt,
		findAllTradeProductsStmt:          q.findAllTradeProductsStmt,
		findClientsByCompanyidStmt:        q.findClientsByCompanyidStmt,
		findOneClientByIdStmt:             q.findOneClientByIdStmt,
		findOneOrderByIdStmt:              q.findOneOrderByIdStmt,
		findPointsBySellerIdStmt:          q.findPointsBySellerIdStmt,
		findProductOrderByIdStmt:          q.findProductOrderByIdStmt,
		getAllProductsStmt:                q.getAllProductsStmt,
		getFornecedorByIDStmt:             q.getFornecedorByIDStmt,
		getProductByIDStmt:                q.getProductByIDStmt,
		updateClientByIdStmt:              q.updateClientByIdStmt,
		updateFornecedorStmt:              q.updateFornecedorStmt,
		updateOneOrderStmt:                q.updateOneOrderStmt,
		updateProductStmt:                 q.updateProductStmt,
		updateProductOrderStmt:            q.updateProductOrderStmt,
		updateTradeProductStmt:            q.updateTradeProductStmt,
	}
}

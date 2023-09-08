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
	if q.createFornecedorStmt, err = db.PrepareContext(ctx, createFornecedor); err != nil {
		return nil, fmt.Errorf("error preparing query CreateFornecedor: %w", err)
	}
	if q.createProductStmt, err = db.PrepareContext(ctx, createProduct); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProduct: %w", err)
	}
	if q.deleteFornecedorByIdStmt, err = db.PrepareContext(ctx, deleteFornecedorById); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteFornecedorById: %w", err)
	}
	if q.deleteProductByIdStmt, err = db.PrepareContext(ctx, deleteProductById); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProductById: %w", err)
	}
	if q.findAllFornecedoresStmt, err = db.PrepareContext(ctx, findAllFornecedores); err != nil {
		return nil, fmt.Errorf("error preparing query FindAllFornecedores: %w", err)
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
	if q.updateFornecedorStmt, err = db.PrepareContext(ctx, updateFornecedor); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateFornecedor: %w", err)
	}
	if q.updateProductStmt, err = db.PrepareContext(ctx, updateProduct); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProduct: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createFornecedorStmt != nil {
		if cerr := q.createFornecedorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createFornecedorStmt: %w", cerr)
		}
	}
	if q.createProductStmt != nil {
		if cerr := q.createProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProductStmt: %w", cerr)
		}
	}
	if q.deleteFornecedorByIdStmt != nil {
		if cerr := q.deleteFornecedorByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteFornecedorByIdStmt: %w", cerr)
		}
	}
	if q.deleteProductByIdStmt != nil {
		if cerr := q.deleteProductByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProductByIdStmt: %w", cerr)
		}
	}
	if q.findAllFornecedoresStmt != nil {
		if cerr := q.findAllFornecedoresStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findAllFornecedoresStmt: %w", cerr)
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
	if q.updateFornecedorStmt != nil {
		if cerr := q.updateFornecedorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateFornecedorStmt: %w", cerr)
		}
	}
	if q.updateProductStmt != nil {
		if cerr := q.updateProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProductStmt: %w", cerr)
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
	db                       DBTX
	tx                       *sql.Tx
	createFornecedorStmt     *sql.Stmt
	createProductStmt        *sql.Stmt
	deleteFornecedorByIdStmt *sql.Stmt
	deleteProductByIdStmt    *sql.Stmt
	findAllFornecedoresStmt  *sql.Stmt
	getAllProductsStmt       *sql.Stmt
	getFornecedorByIDStmt    *sql.Stmt
	getProductByIDStmt       *sql.Stmt
	updateFornecedorStmt     *sql.Stmt
	updateProductStmt        *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                       tx,
		tx:                       tx,
		createFornecedorStmt:     q.createFornecedorStmt,
		createProductStmt:        q.createProductStmt,
		deleteFornecedorByIdStmt: q.deleteFornecedorByIdStmt,
		deleteProductByIdStmt:    q.deleteProductByIdStmt,
		findAllFornecedoresStmt:  q.findAllFornecedoresStmt,
		getAllProductsStmt:       q.getAllProductsStmt,
		getFornecedorByIDStmt:    q.getFornecedorByIDStmt,
		getProductByIDStmt:       q.getProductByIDStmt,
		updateFornecedorStmt:     q.updateFornecedorStmt,
		updateProductStmt:        q.updateProductStmt,
	}
}

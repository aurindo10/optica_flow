package database

import (
	"database/sql"
	"fmt"
	repository "optica_flow/internal/app/infra/database/productRepository"
	database "optica_flow/internal/app/infra/database/queries"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose"
	"go.uber.org/fx"
)



func NewPostgresDatabase() (*database.Queries, error ){
	db, err := sql.Open("postgres", "user=docker dbname=otica password=otica123 sslmode=disable host=localhost port=5432")
	if err != nil {
		return nil , err
	}
	queries := database.New(db)
	return queries, nil
}

func NewPostgresTestDatabase() (*database.Queries, func(), error) {
	// Gerar um nome de banco de dados único para evitar colisões
	tempDBName := fmt.Sprintf("test_db_%d", time.Now().Unix())

	// Conectar ao banco de dados principal para criar um novo banco de dados temporário
	masterDB, err := sql.Open("postgres", "user=docker dbname=otica password=otica123 sslmode=disable host=localhost port=5432")
	if err != nil {
		return nil, nil, err
	}
	_, err = masterDB.Exec(fmt.Sprintf("CREATE DATABASE %s;", tempDBName))
	if err != nil {
		return nil, nil, err
	}

	// Conectar ao novo banco de dados temporário
	tempDB, err := sql.Open("postgres", fmt.Sprintf("user=docker dbname=%s password=otica123 sslmode=disable host=localhost port=5432", tempDBName))
	if err != nil {
		return nil, nil, err
	}

	// Executar migrações Goose
	if err := goose.Up(tempDB, "../../database/sql/migrations"); err != nil {
		return nil, nil, fmt.Errorf("falha na migração: %w", err)
	}

	queries := database.New(tempDB)

	// Função para limpar o banco de dados temporário
	cleanup := func() {
		tempDB.Close()
		masterDB.Exec(fmt.Sprintf("DROP DATABASE %s;", tempDBName))
		masterDB.Close()
	}

	return queries, cleanup, nil
}


var TestModule = fx.Options(
	fx.Provide(
		NewPostgresTestDatabase,
	),
	repository.Module,)


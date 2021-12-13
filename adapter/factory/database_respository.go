package factory

import (
	"database/sql"

	repo "github.com/JPauloMoura/payment-gateway/adapter/repository"
	"github.com/JPauloMoura/payment-gateway/domain/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repo.NewTransactionRepositoryDB(r.DB)
}

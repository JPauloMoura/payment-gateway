package factory

import "github.com/JPauloMoura/payment-gateway/domain/repository"

// RepositoryFactory é a inteface que define as regras para criação de um repository.TransactionRepository
type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}

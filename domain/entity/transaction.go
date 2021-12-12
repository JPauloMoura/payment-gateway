package entity

import (
	"context"
	"log"
)

// Transaction é a entidade padrão responsavel por uma transação
type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

// NewTransaction retorna uma nova entidade Transaction vazia
func NewTransaction() *Transaction {
	return &Transaction{}
}

// IsValid verifica se a Transaction criada pode ser realizada
func (t Transaction) IsValid() error {
	if t.Amount > 1000 || t.Amount < 1 {
		log.Printf("context: %v, err: %s", context.Background(), ErrLimitForTransaction.Error())
		return ErrLimitForTransaction
	}

	return nil
}

// SetCreditCard adiciona um cartão de credito a transação
func (t *Transaction) SetCreditCard(card CreditCard) {
	t.CreditCard = card
}

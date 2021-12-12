package fixture

import (
	"context"
	"database/sql"
	"log"
	"time"
)

// TransactionRepositoryDB é responsavel por manipular os dados de transação em um BD
type TransactionRepositoryDB struct {
	db *sql.DB
}

// NewTransactionRepositoryDB retorna uma nova struct TransactionRepositoryDB
// que permite manipular os dados de transação em um BD
func NewTransactionRepositoryDB(db *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{db: db}
}

// Insert é responsável por persistir os dados de uma transação
func (t *TransactionRepositoryDB) Insert(id string, account string, amount float64, status string, errorMessage string) error {
	// prepara o script de insersão no BD
	stmt, err := t.db.Prepare(`
		INSERT INTO transactions (id, account_id, amount, status, error_message, created_at, updated_at)
		values($1,$2,$3,$4,$5,$6,$7)
	`)

	if err != nil {
		log.Printf("context: %v, err: %s", context.Background(), err.Error())
		return err
	}

	// executa o script passando os campos necessários
	_, err = stmt.Exec(
		id,
		account,
		amount,
		status,
		errorMessage,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		log.Printf("context: %v, err: %s", context.Background(), err.Error())
		return err
	}
	return nil
}

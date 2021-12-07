package repository

// TransactionRepository é a interface que define o contrato para realizar uma transação
type TransactionRepository interface {
	Insert(id string, account string, amount float64, status string, errorMessage string) error
}

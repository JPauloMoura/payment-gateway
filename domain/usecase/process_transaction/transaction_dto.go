package process_transaction

/*
	Aqui definimos o padrão a entrada e saida da nossa transação
*/

// TransactionDtoInput define os dados que devem ser utilizados como input de uma transação
type TransactionDtoInput struct {
	ID                        string  `json:"id"`
	AccountID                 string  `json:"account_id"`
	CreditCardNumber          string  `json:"credit_card_number"`
	CreditCardName            string  `json:"credit_card_name"`
	CreditCardExpirationMonth int     `json:"credit_card_expiration_month"`
	CreditCardExpirationYear  int     `json:"credit_card_expiration_year"`
	CreditCardCVV             int     `json:"credit_card_cvv"`
	Amount                    float64 `json:"amount"`
}

// TransactionDtoOutput define os dados que devem ser utilizados como output de uma transação
type TransactionDtoOutput struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

package transaction

import (
	"encoding/json"

	"github.com/JPauloMoura/payment-gateway/domain/usecase/process_transaction"
)

// KafkaPresenter é contém os dados que devem ser enviados para o kafra
type KafkaPresenter struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

// NewTransactionKafkaPresenter retorna uma nova struct KafkaPresenter vazia
func NewTransactionKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

// Bind implementa a interface Presenter e recebe os dados de output do usecase e os insere
// em uma estrutura KafkaPresenter
func (t *KafkaPresenter) Bind(input interface{}) error {
	t.ID = input.(process_transaction.TransactionDtoOutput).ID
	t.Status = input.(process_transaction.TransactionDtoOutput).Status
	t.ErrorMessage = input.(process_transaction.TransactionDtoOutput).ErrorMessage
	return nil
}

// Show implementa a interface Presenter e retorna os dados do KafkaPresenter ou um erro
func (t *KafkaPresenter) Show() (jsonData []byte, err error) {
	jsonData, err = json.Marshal(t)
	if err != nil {
		return
	}

	return
}

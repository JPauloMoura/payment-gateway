package process_transaction

import (
	"github.com/JPauloMoura/payment-gateway/adapter/broker"
	"github.com/JPauloMoura/payment-gateway/domain/entity"
	"github.com/JPauloMoura/payment-gateway/domain/repository"
)

// ProcessTransaction define a estrutura de dados para realizar o processamento de uma trasação
type ProcessTransaction struct {
	Repository repository.TransactionRepository
	Producer   broker.Producer
	Topic      string
}

// NewProcessTransaction devolve uma processo de transação que pode ser realizado
func NewProcessTransaction(repository repository.TransactionRepository, producer broker.Producer, topic string) *ProcessTransaction {
	return &ProcessTransaction{
		Repository: repository,
		Producer:   producer,
		Topic:      topic,
	}
}

// Execute realiza a execução de um processo de transação
func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	//cria uma nova transaction
	transaction := entity.NewTransaction()

	// seta seus dados
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	// cria um cartão de credito com os dados inputados
	cc, err := entity.NewCrediCard(
		input.CreditCardNumber,
		input.CreditCardName,
		input.CreditCardExpirationMonth,
		input.CreditCardExpirationYear,
		input.CreditCardCVV,
	)

	if err != nil {
		return p.rejectTransaction(transaction, err)
	}

	transaction.SetCreditCard(*cc)

	if err := transaction.IsValid(); err != nil {
		return p.rejectTransaction(transaction, err)
	}

	return p.approveTransaction(transaction)
}

// approveTransaction se a transação for aprovada os dados devem ser inseridos com status: "approveded"
func (p *ProcessTransaction) approveTransaction(transaction *entity.Transaction) (TransactionDtoOutput, error) {

	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.APPROVED, "")
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	// publica o evento de aprovação
	if err = p.publish(output, []byte(transaction.ID)); err != nil {
		return TransactionDtoOutput{}, err
	}

	return output, nil
}

// rejectTransaction se a transação não for aprovada os dados devem ser inseridos com status: "rejected"
func (p ProcessTransaction) rejectTransaction(transaction *entity.Transaction, e error) (TransactionDtoOutput, error) {

	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.REJECTED, e.Error())
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       entity.REJECTED,
		ErrorMessage: e.Error(),
	}

	// publica o evento de rejeição
	if err = p.publish(output, []byte(transaction.ID)); err != nil {
		return TransactionDtoOutput{}, err
	}

	return output, nil
}

// publish realiza a publicação do evento no Producer que foi implementado
func (p *ProcessTransaction) publish(output TransactionDtoOutput, key []byte) error {
	err := p.Producer.Publish(output, key, p.Topic)
	if err != nil {
		return err
	}
	return nil
}

package process_transaction

import (
	"testing"
	"time"

	mock_producer "github.com/JPauloMoura/payment-gateway/adapter/broker/mock"
	"github.com/JPauloMoura/payment-gateway/domain/entity"
	mock_repository "github.com/JPauloMoura/payment-gateway/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// Cernário 1:
// A transação não pode ser executada pois
// o cartão de credito é inválido, de acordo com as regras definidas para ele.
func TestProcessTransaction_Execute_InvalidCrediCart(t *testing.T) {

	// define qual será a entrada para nossa transação
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "40000000000000000", // número inválido
		CreditCardName:            "José da Silva Pereira",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}

	// define qual será a saída esperada para nossa transação
	expected := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: entity.ErrInvalidCreditCardNumber.Error(),
	}

	// realiza o controle do nosso mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// cria um novo mock de um repository e de um producer
	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	producerMock := mock_producer.NewMockProducer(ctrl)

	// realiza o mock da entrada e saida do repository e producer
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expected.Status, expected.ErrorMessage).
		Return(nil)
	producerMock.EXPECT().Publish(expected, []byte(input.ID), "transactions_result")

	processTransaction := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")
	result, err := processTransaction.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

// Cernário 2:
// A transação não pode ser executada pois
// a transação é inválida, de acordo com as regras definidas para ela.
func TestProcessTransaction_Execute_InvalidTransaction(t *testing.T) {

	// define qual será a entrada para nossa transação
	input := TransactionDtoInput{
		ID:                        "123456",
		AccountID:                 "0123",
		CreditCardNumber:          "5203902335999315",
		CreditCardName:            "José da Silva Pereira",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             608,
		Amount:                    0, // valor inválida oara transação
	}

	// define qual será a saída esperada para nossa transação
	expected := TransactionDtoOutput{
		ID:           "123456",
		Status:       entity.REJECTED,
		ErrorMessage: entity.ErrLimitForTransaction.Error(),
	}

	// realiza o controle do nosso mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// cria um novo mock de um repository e de um producer
	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	producerMock := mock_producer.NewMockProducer(ctrl)

	// realiza o mock da entrada e saida do repository e producer
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expected.Status, expected.ErrorMessage).
		Return(nil)
	producerMock.EXPECT().Publish(expected, []byte(input.ID), "transactions_result")

	processTransaction := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")
	result, err := processTransaction.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

// Cernário 3:
// A transação é realizado com sucesso pois:
// a transação e o cartão de crédito são válidos, de acordo com suas regras.
func TestProcessTransaction_Execute_ApprovedTransaction(t *testing.T) {

	// define qual será a entrada para nossa transação
	input := TransactionDtoInput{
		ID:                        "123456",
		AccountID:                 "0123",
		CreditCardNumber:          "5203902335999315",
		CreditCardName:            "José da Silva Pereira",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             608,
		Amount:                    100,
	}

	// define qual será a saída esperada para nossa transação
	expected := TransactionDtoOutput{
		ID:           "123456",
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	// realiza o controle do nosso mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// cria um novo mock de um repository e de um producer
	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	producerMock := mock_producer.NewMockProducer(ctrl)

	// realiza o mock da entrada e saida do repository e producer
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expected.Status, expected.ErrorMessage).
		Return(nil)
	producerMock.EXPECT().Publish(expected, []byte(input.ID), "transactions_result")

	processTransaction := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")
	result, err := processTransaction.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

package kafka

import (
	"testing"

	"github.com/JPauloMoura/payment-gateway/adapter/presenter/transaction"
	"github.com/JPauloMoura/payment-gateway/domain/entity"
	"github.com/JPauloMoura/payment-gateway/domain/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

// Cenário 1:
// publicada uma msg no kafka com sucesso
func TestProducerPublish_Sucess(t *testing.T) {
	expected := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: entity.ErrLimitForTransaction.Error(),
	}

	// essa configuração diz para o kafka que estamos realizando um teste
	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}

	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expected, []byte("1"), "test")
	assert.Nil(t, err)
}

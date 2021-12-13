package kafka

import (
	"github.com/JPauloMoura/payment-gateway/adapter/presenter"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// Producer é a struct responsável por produzir/emitir os eventos para o kafka
type Producer struct {
	ConfigMap *ckafka.ConfigMap
	Presenter presenter.Presenter
}

// NewKafkaProducer retorna uma novo Producer com as configurações definidas
func NewKafkaProducer(configMap *ckafka.ConfigMap, presenter presenter.Presenter) *Producer {
	return &Producer{ConfigMap: configMap, Presenter: presenter}
}

// Publish é o metódo responsável por emitir os eventos para o kafka.
// Ele recebe uma mensagem, a key da partição, e o nome do topico
func (p *Producer) Publish(msg interface{}, key []byte, topic string) error {
	producer, err := ckafka.NewProducer(p.ConfigMap)
	if err != nil {
		return err
	}

	if err := p.Presenter.Bind(msg); err != nil {
		return err
	}

	presenterMsg, err := p.Presenter.Show()
	if err != nil {
		return err
	}

	// configura a msg no kafka
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          presenterMsg,
		Key:            key,
	}

	// emite a msg

	if err := producer.Produce(message, nil); err != nil {
		panic(err)
	}

	return nil
}

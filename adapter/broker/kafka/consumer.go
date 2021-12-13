package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

// Consumer é a struct responsável por consumir os eventos emitidos pelo kafka
type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

// NewConsumer retorna uma struct Consumer configurada
func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

// Consume recebe um *channel e fica ouvindo as msg que chegam nos tópicos
// que ele estiver inscrito e adiciona essas mensagens dentro do *channel
func (c *Consumer) Listen(channel chan *ckafka.Message) error {
	// cria um novo consumer
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}

	// inscreve esse consumer para ficar consumindo msg de um topico

	if err := consumer.SubscribeTopics(c.Topics, nil); err != nil {
		panic(err)
	}

	// loop para ficar ouvindo as msgs
	for {
		msg, err := consumer.ReadMessage(-1)

		// se não tive error, essa msg e jogada dentro de um channel
		if err == nil {
			channel <- msg
		}
	}
}

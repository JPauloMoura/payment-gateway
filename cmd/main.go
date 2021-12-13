package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/JPauloMoura/payment-gateway/adapter/broker/kafka"
	"github.com/JPauloMoura/payment-gateway/adapter/factory"
	"github.com/JPauloMoura/payment-gateway/adapter/presenter/transaction"
	"github.com/JPauloMoura/payment-gateway/domain/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// configuração do BD
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	// criação do repository
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	// configuração do producer
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	// configuração do consumer
	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}

	// definindo todos os topicos
	topics := []string{"transactions"}

	// cria o consumer
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	// as messagens vão ficar sendo lidas em uma thred sepada
	go consumer.Listen(msgChan)

	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	// lemos todas messagens que entram no channel msgChan
	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}

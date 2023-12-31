package kafka

import (
	"fmt"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	"github.com/rodolfoHOk/fullcycle.imersao15/tree/main/codepix/application/factory"
	appModel "github.com/rodolfoHOk/fullcycle.imersao15/tree/main/codepix/application/model"
	"github.com/rodolfoHOk/fullcycle.imersao15/tree/main/codepix/application/usecase"
	"github.com/rodolfoHOk/fullcycle.imersao15/tree/main/codepix/domain/model"
)

type KafkaProcessor struct {
	Database *gorm.DB
	Producer *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

func NewKafkaProcessor(database *gorm.DB, producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Database: database,
		Producer: producer,
		DeliveryChan: deliveryChan,
	}
}

func (processor *KafkaProcessor) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id": os.Getenv("kafkaConsumerGroupId"),
		"auto.offset.reset":"earliest",
		"security.protocol": os.Getenv("security.protocol"),
		"sasl.mechanisms": os.Getenv("sasl.mechanisms"),
		"sasl.username": os.Getenv("sasl.username"),
		"sasl.password": os.Getenv("sasl.password"),
	}
	consumer, err := ckafka.NewConsumer(configMap);
	if err != nil {
		panic(err)
	}

	topics := []string{os.Getenv("kafkaTransactionTopic"), os.Getenv("kafkaTransactionConfirmationTopic")}
	consumer.SubscribeTopics(topics, nil)
	fmt.Println("kafka consumer has been started")

	for {
		kafkaMessage, err := consumer.ReadMessage(-1)
		if err == nil {
			processor.processMessage(kafkaMessage)
		}
	}
}

func (processor *KafkaProcessor) processMessage(kafkaMessage *ckafka.Message) {
	transactionsTopic := "transactions";
	transactionConfirmationTopic := "transaction_confirmation";

	topic := kafkaMessage.TopicPartition.Topic;
	if (*topic == transactionsTopic) {
		processor.processTransaction(kafkaMessage)
	} else if (*topic == transactionConfirmationTopic) {
		processor.processTransactionConfirmation(kafkaMessage)
	} else {
		fmt.Println("not a valid topic", string(kafkaMessage.Value))
	}
}

func (processor *KafkaProcessor) processTransaction(kafkaMessage *ckafka.Message) error {
	transaction := appModel.NewTransaction()
	err := transaction.ParseJson(kafkaMessage.Value)
	if err != nil {
		return err
	}

	transactionUseCase := factory.TransactionUseCaseFactory(processor.Database)

	createdTransaction, err := transactionUseCase.Register(
		transaction.AccountID,
		transaction.Amount,
		transaction.PixKeyTo,
		transaction.PixKeyKindTo,
		transaction.Description,
		transaction.ID,
	)
	if err != nil {
		fmt.Println("error registering transaction", err)
		return err
	}

	topic := "bank" + createdTransaction.PixKeyTo.Account.Bank.Code
	transaction.ID = createdTransaction.ID
	transaction.Status = model.TransactionPending
	transactionJson, err := transaction.ToJson()
	if err != nil {
		return err
	}

	err = Publish(string(transactionJson), topic, processor.Producer, processor.DeliveryChan)
	if err != nil {
		return err
	}

	return nil
}

func (processor *KafkaProcessor) processTransactionConfirmation(kafkaMessage *ckafka.Message) error {
	transaction := appModel.NewTransaction()
	err := transaction.ParseJson(kafkaMessage.Value)
	if err != nil {
		return err
	}
	println(transaction.ID)
	println(string(kafkaMessage.Value))

	transactionUseCase := factory.TransactionUseCaseFactory(processor.Database)

	if transaction.Status == model.TransactionConfirmed {
		err = processor.confirmTransaction(transaction, transactionUseCase)
		if err != nil {
			return err
		}
	} else if transaction.Status == model.TransactionCompleted {
		_, err := transactionUseCase.Complete(transaction.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (processor *KafkaProcessor) confirmTransaction(transaction *appModel.Transaction, transactionUseCase usecase.TransactionUseCase) error {
	confirmedTransaction, err := transactionUseCase.Confirm(transaction.ID)
	if err != nil {
		return err
	}

	topic := "bank" + confirmedTransaction.AccountFrom.Bank.Code
	transactionJson, err := transaction.ToJson()
	if err != nil {
		return err
	}

	err = Publish(string(transactionJson), topic, processor.Producer, processor.DeliveryChan)
	if err != nil {
		return err
	}

	return nil
}

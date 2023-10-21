package kafka

import (
	"fmt"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"security.protocol": os.Getenv("security.protocol"),
		"sasl.mechanisms":   os.Getenv("sasl.mechanisms"),
		"sasl.username":     os.Getenv("sasl.username"),
		"sasl.password":     os.Getenv("sasl.password"),
	}
	producer, err := ckafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}
	return producer
}

func Publish(message string, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
	kafkaMessage := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic: &topic,
			Partition: ckafka.PartitionAny,
		},
		Value: []byte(message),
	} 
	err := producer.Produce(kafkaMessage, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan ckafka.Event) {
	for kafkaEvent := range deliveryChan {
		switch event := kafkaEvent.(type) {
		case *ckafka.Message: 
			if event.TopicPartition.Error != nil {
				fmt.Println("Delivery failed:", event.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", event.TopicPartition)
			}
		}
	}
}

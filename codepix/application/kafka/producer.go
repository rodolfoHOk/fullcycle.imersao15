package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	producer, err := ckafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}
	return producer
}

func Publish(message string, topic string, producer *ckafka.Producer) error {
	kafkaMessage := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic: &topic,
			Partition: ckafka.PartitionAny,
		},
		Value: []byte(message),
	} 
	err := producer.Produce(kafkaMessage, nil)
	if err != nil {
		return err
	}
	return nil
}

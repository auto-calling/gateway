package handler

import (
	"github.com/auto-calling/gateway/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

func Producer(msg []byte) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaBrokerList,
	})
	if err != nil {
		return err
	}
	defer p.Close()
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Error("Delivery failed: %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	// Produce messages to topic (asynchronously)
	topic := config.KafkaTopic
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
	return err
}

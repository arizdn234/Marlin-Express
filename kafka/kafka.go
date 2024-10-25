package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

// Kafka Producer
func NewKafkaWriter(brokerURL, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{brokerURL},
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: 10 * time.Second,
	})
}

// Kafka Consumer
func NewKafkaReader(brokerURL, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerURL},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

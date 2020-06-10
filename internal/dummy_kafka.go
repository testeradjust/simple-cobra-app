package internal

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type DummyKafka struct {
	producer sarama.AsyncProducer
}

func NewDummyKafka(brokers []string) (*DummyKafka, error) {
	producer, err := newAsyncProducer(brokers)
	if err != nil {
		return nil, err
	}

	return &DummyKafka{
		producer: producer,
	}, nil
}

func newAsyncProducer(brokerList []string) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		return nil, fmt.Errorf("failed to start kafka producer: %s", err)
	}

	return producer, nil
}

func (dk *DummyKafka) ProduceTestMessage(message string) {
	dk.producer.Input() <- &sarama.ProducerMessage{
		Topic: "dummy-topic",
		Key:   sarama.StringEncoder("dummy-key"),
		Value: sarama.StringEncoder(message),
	}
}

func (dk *DummyKafka) Close() error {
	if err := dk.producer.Close(); err != nil {
		return fmt.Errorf("failed to shut down async producer: %s", err)
	}
	return nil
}

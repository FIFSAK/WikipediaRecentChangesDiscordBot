package kafka

import (
	"WikipediaRecentChangesDiscordBot/config"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type Kafka struct {
	Config  *config.Config
	Context context.Context
	Writer  *kafka.Writer
	Reader  *kafka.Reader
}

func NewKafka(config *config.Config) *Kafka {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(config.KafkaBroker),
		Topic:    config.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.KafkaBroker},
		Topic:   config.KafkaTopic,
		GroupID: config.KafkaGroup,
	})
	return &Kafka{
		Config:  config,
		Context: context.Background(),
		Writer:  writer,
		Reader:  reader,
	}
}

func (k *Kafka) SendKafka(date string, language string) error {
	event := fmt.Sprintf("%s:%s", date, language)

	err := k.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(event),
		},
	)
	if err != nil {
		log.Fatalf("failed to write message: %v", err)
		return err
	}

	return nil
}

func (k *Kafka) Close() error {
	if err := k.Writer.Close(); err != nil {
		return err
	}
	if err := k.Reader.Close(); err != nil {
		return err
	}
	return nil
}

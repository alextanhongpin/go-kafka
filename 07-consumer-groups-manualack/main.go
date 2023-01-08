package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	var topic string
	flag.StringVar(&topic, "topic", "topic", "set the kafka topic")
	flag.Parse()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9093"},
		GroupID:  "consumer-group-id",
		Topic:    topic,
		MinBytes: 10e3, //10KB
		MaxBytes: 10e4, // 10MB
	})

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

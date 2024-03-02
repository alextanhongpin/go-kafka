package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	var role, topic string
	flag.StringVar(&topic, "topic", "topic", "set the kafka topic")
	flag.StringVar(&role, "role", "writer", "reader or writer")
	flag.Parse()

	switch role {
	case "reader":
		read(topic)
	case "writer":
		if err := write(topic); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown role: %s", role)
	}
}

func write(topic string) error {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9093"),
		Topic:    topic,
		Balancer: &kafka.Hash{}, // Default is round robin.
	}

	// The first message is a poison pill.
	// This message will fail to be processed by the
	// consumer, and blocks future messages.
	// So the offset will always remain at the last offset, causing consumer lag.
	if err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("a"),
			Value: []byte("poison"),
		},
		kafka.Message{
			Key:   []byte("a"),
			Value: []byte("a is for alice"),
		},
		kafka.Message{
			Key:   []byte("b"),
			Value: []byte("b is for bob"),
		},
		kafka.Message{
			Key:   []byte("c"),
			Value: []byte("c is for charles"),
		},
	); err != nil {
		return fmt.Errorf("failed to write messages: %w", err)
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	return nil
}

func read(topic string) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9093"},
		GroupID:     "consumer-group-id",
		Topic:       topic,
		MinBytes:    10e3, //10KB
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.LastOffset,
	})
	ctx := context.Background()

	var m kafka.Message
	var err error

	defer func() {
		// If the system panic due to poison pill, we recover and send the poison
		// pill to the poison pill topic.
		if rec := recover(); rec != nil {
			w := &kafka.Writer{
				Addr:     kafka.TCP("localhost:9093"),
				Topic:    "poison-pill",
				Balancer: &kafka.Hash{}, // Default is round robin.
			}

			writeErr := w.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte("a"),
					Value: []byte("poison"),
				},
			)
			if writeErr != nil {
				err = fmt.Errorf("%w: %w", err, writeErr)
				return
			}

			// Commits the offset of the poison pill.
			if commitErr := r.CommitMessages(ctx, m); commitErr != nil {
				err = fmt.Errorf("%w: %w", err, commitErr)
				return
			}
		}
	}()

	for {
		m, err = r.FetchMessage(ctx)
		if err != nil {
			// When we reach the end of the topic, we break the loop.
			err = fmt.Errorf("failed to fetch message: %w", err)
			break
		}

		// If we detect a poison pill, we panic.
		// We can write a pattern matching to detect the poison pill too.
		if string(m.Value) == "poison" {
			panic("poison pill")
			// This will cause the consumer to restart.
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		if err := r.CommitMessages(ctx, m); err != nil {
			return fmt.Errorf("failed to commit messages: %w", err)
		}
	}

	return err
}

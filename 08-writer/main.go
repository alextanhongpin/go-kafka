package main

import (
	"context"
	"flag"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	var topic string
	flag.StringVar(&topic, "topic", "topic", "set the kafka topic")
	flag.Parse()

	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9093"),
		Topic:    topic,
		Balancer: &kafka.Hash{}, // Default is round robin.
	}

	if err := w.WriteMessages(context.Background(),
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
		// This will be in the same partition as "c is for charles"
		kafka.Message{
			Key:   []byte("c"),
			Value: []byte("c is for cat"),
		},
		kafka.Message{
			Key:   []byte("d"),
			Value: []byte("d is for damian"),
		},
		kafka.Message{
			Key:   []byte("e"),
			Value: []byte("e is for elise"),
		},
	); err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

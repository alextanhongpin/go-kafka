package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	var topic string
	var partition int
	flag.StringVar(&topic, "topic", "topic", "set a topic")
	flag.IntVar(&partition, "partition", 0, "set a partition")
	flag.Parse()

	ctx := context.Background()
	conn, err := kafka.DialLeader(ctx, "tcp", "localhost:9093", topic, partition)
	if err != nil {
		log.Fatalf("failed to dial leader: %v", err)
	}

	// NOTE: This only writes to one partition.
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one")},
		kafka.Message{Value: []byte("two")},
		kafka.Message{Value: []byte("three")},
	)
	if err != nil {
		log.Fatalf("failed to write messages: %v", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatalf("failed to close writer: %v", err)
	}
}

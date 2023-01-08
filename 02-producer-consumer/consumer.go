package main

import (
	"context"
	"flag"
	"fmt"
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

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	minBytes := int(10e3)                       // 10KB
	maxBytes := int(1e6)                        // 1MB
	batch := conn.ReadBatch(minBytes, maxBytes) // Fetch 10KB min, 1MB max.
	b := make([]byte, minBytes)                 // 10KB max per message.
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}
	if err := batch.Close(); err != nil {
		log.Fatalf("failed to close batch: %v", err)
	}
	if err := conn.Close(); err != nil {
		log.Fatalf("failed to close writer: %v", err)
	}
}

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
	var partition int
	var offset int64
	flag.StringVar(&topic, "topic", "topic", "set a topic")
	flag.IntVar(&partition, "partition", 0, "set a partition")
	flag.Int64Var(&offset, "offset", 0, "set a offset")
	flag.Parse()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9093"},
		Topic:     topic,
		Partition: partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(offset)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

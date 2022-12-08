package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9093", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	h := &kafka.Hash{}
	for i := 0; i < 10; i++ {
		fmt.Println(i, h.Balance(kafka.Message{Key: []byte(fmt.Sprint(i)), Value: []byte("hello world")}, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	}

	if err := write(conn, topic); err != nil {
		log.Fatal("failed to write message:", err)
	}
	if err := read(conn); err != nil {
		log.Fatal("failed to read message:", err)
	}
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func write(conn *kafka.Conn, topic string) error {
	w := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9093"),
		Topic:        topic,
		WriteTimeout: 10 * time.Second,
		Balancer:     &kafka.Hash{}, // Default is round-robin.
	}
	defer w.Close()
	if err := w.WriteMessages(
		context.Background(),
		kafka.Message{Key: []byte("1"), Value: []byte("one!")},
		kafka.Message{Key: []byte("4"), Value: []byte("two!")},
		kafka.Message{Key: []byte("4"), Value: []byte("three")},
	); err != nil {
		return err
	}
	return nil
}

func read(conn *kafka.Conn) error {
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // Fetch 10kb min, 1MB max.

	b := make([]byte, 10e3) // 10kb max message.
	for {
		n, err := batch.Read(b)
		if err != nil {
			return err
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		return fmt.Errorf("failed to close batch: %w", err)
	}

	return nil
}

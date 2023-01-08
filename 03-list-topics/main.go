package main

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.Dial("tcp", "localhost:9093")
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		log.Fatalf("failed to read partitions: %v", err)
	}

	for _, p := range partitions {
		fmt.Printf("%+v\n", p)
	}
}

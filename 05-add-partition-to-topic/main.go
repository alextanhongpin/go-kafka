package main

import (
	"context"
	"flag"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	var topic string
	var partition int
	flag.IntVar(&partition, "partition", 3, "set the number of partitions. Note that the partition count cannot be lower than existing partition")
	flag.StringVar(&topic, "topic", "topic", "set the kafka topic")
	flag.Parse()

	client := &kafka.Client{
		Addr: kafka.TCP("localhost:9093"),
	}
	res, err := client.CreatePartitions(context.Background(), &kafka.CreatePartitionsRequest{
		Topics: []kafka.TopicPartitionsConfig{
			{
				Name:  topic,
				Count: int32(partition),
			},
		},
	})
	if err != nil {
		log.Fatal("failed to create/update topic:", err)
	}

	log.Printf("%+v\n", res)
}

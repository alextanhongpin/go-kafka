package main

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	// This example shows how the partition is determined from the message `Key`.
	// We have partitions numbered 1 to 5.
	// We pass the key a to z and see the index of the partitions generated from
	// 0 to 4.
	// Running this multiple times will always produce the same result.
	h := &kafka.Hash{}
	partitions := []int{1, 2, 3, 4, 5}

	for i := 0; i < 26; i++ {
		msg := kafka.Message{
			Key: []byte(string('a' + rune(i))),
		}
		fmt.Println(string('a'+rune(i)), h.Balance(msg, partitions...))
	}
}

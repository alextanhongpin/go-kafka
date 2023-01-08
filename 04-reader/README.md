# 04-reader
This example demonstrates reading the message from a given topic and partition, and at the given offset.


Try running the example with different offset after producing some messages:


```bash
# Run this multiple times to produce messages.
$ go run 02-producer-consumer/producer.go

# Run the reader at different offset to read messages.
$ go run 04-reader/main.go -offset 3
$ go run 04-reader/main.go -offset 10
```

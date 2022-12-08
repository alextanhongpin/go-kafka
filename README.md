# go-kafka


## Setup

Start the docker containers:
```bash
$ make up
```

Access the kafka cli:
```bash
$ docker exec -it $(docker ps --filter name=kafka-1 -q) bash

# In the container, cd to the directory where the bash files are located.
$ cd opt/bitnami/kafka
```

## Playing with kafka

```bash
# Creating topics
$ bin/kafka-topics.sh --create --topic my-topic --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1

# Viewing topics
$ bin/kafka-topics.sh --bootstrap-server=localhost:9092 --describe --topic my-topic

# List all topics
$ bin/kafka-topics.sh --bootstrap-server=localhost:9092 --list

# Product events.
$ bin/kafka-console-producer.sh --topic my-topic --bootstrap-server localhost:9092

# View events.
$ bin/kafka-console-consumer.sh --topic my-topic --from-beginning --bootstrap-server localhost:9092

# View events from specific partition.
$ bin/kafka-console-consumer.sh --topic my-topic --from-beginning --bootstrap-server localhost:9092 --partition 0

# Add partitions to topics.
$ bin/kafka-topics.sh --bootstrap-server=localhost:9092 --alter --topic my-topic --partitions 10

# Check distribution across all partitions.
$ bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list localhost:9092 --topic my-topic --time -1
```

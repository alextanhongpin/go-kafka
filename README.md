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


# View events from a given offset
$ bin/kafka-console-consumer.sh --topic topic -bootstrap-server localhost:9092 --offset 12 --partition 0

# Add partitions to topics.
$ bin/kafka-topics.sh --bootstrap-server=localhost:9092 --alter --topic my-topic --partitions 10

# Check distribution across all partitions.
$ bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list localhost:9092 --topic my-topic --time -1

# View consumer group offset
$ bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group consumer-group-id --describe
```

## Questions
- how to add partitions to existing topic?

```bash
# Alter partition
$ bin/kafka-topics.sh --bootstrap-server localhost:9092 --alter --topic topic --partitions 40
```

```
# Verify (view partitions)
```

Output:

```bash
$ bin/kafka-topics.sh --bootstrap-server=localhost:9092 --describe --topic topic
Topic: topic    TopicId: ya8FFGKWR5-h8P0jWn5tGw PartitionCount: 3       ReplicationFactor: 1    Configs:
        Topic: topic    Partition: 0    Leader: 1001    Replicas: 1001  Isr: 1001
        Topic: topic    Partition: 1    Leader: 1001    Replicas: 1001  Isr: 1001
        Topic: topic    Partition: 2    Leader: 1001    Replicas: 1001  Isr: 1001
```

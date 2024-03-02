# Poison Pill

## What

A poison pill (in the context of Kafka) is a record that has been produced to a Kafka topic and always fails when consumed, no matter how many times it is attempted[^1].

[^1]: https://www.confluent.io/blog/spring-kafka-can-your-kafka-consumers-handle-a-poison-pill/


## How do we simulate this?


For simplicity, we publish a numeric string as the message. If the message is `"10"`, then we treat it a poison pill message.

This message cannot be completed, so we will publish it into another message queue to be handled by another service.


In short, we need the following capability:
- the ability to differentiate which message is poison pill (we can use some rule engine, or evaluate the message payload). This can be configured through some config or environment variables, so that we can restart the consumer to skip this message.
- we need another queue to publish the poison pill message. Some kind of monitoring tools is necessary to observe the poison pill message.
- we need one service to handle the poison pill message. Since this is manual, we can use some kind of UI to handle this message. We can also publish the message back to the original queue after the issue is resolved. In short, we need to know the original queue too in order to publish it back.



## Steps


1. create two topics: `topic` and `poison-pill`
2. start the kafka and zookeeper
3. create a topic
4. publish the message to the topic
5. consume the message

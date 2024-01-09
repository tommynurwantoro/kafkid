# :bullettrain_front: Kafkid - Kafka Messaging Service

This project provides a simple messaging service that allows publishing and consuming messages through Kafka. It includes an API for message publishing and a consumer for message consumption.

## Feature

- Publish messages to Kafka through a simple API.
- Consume messages from Kafka using a consumer.
- Configurable settings for Kafka connection and topics.

## Prerequisites

Make sure you have the following installed before setting up and running the messaging service:

- Kafka
- Zookeeper

This project provides `zk-single-kafka-single.yaml` you can run this command to easy setup Kafka and Zookeeper using docker

```bash
docker-compose -f zk-single-kafka-single.yaml up -d
```

## Getting Started

### Configuration

Copy `config.sample.yaml` to `config.yaml` and modify based on your needs.

### Download Package

```bash
make install
```


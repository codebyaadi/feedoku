#!/bin/bash

# Script to start Apache Kafka using Docker

# Set environment variables
ZOOKEEPER_CONTAINER_NAME="zookeeper"
KAFKA_CONTAINER_NAME="kafka"
NETWORK_NAME="kafka-network"
ZOOKEEPER_PORT=2181
KAFKA_PORT=9092

# Check if the Docker network exists; if so, remove it
if docker network ls --format '{{.Name}}' | grep -q "^${NETWORK_NAME}$"; then
  echo "Docker network $NETWORK_NAME already exists. Removing it..."
  docker network rm $NETWORK_NAME
fi

# Create a custom network for Kafka and Zookeeper to communicate
echo "Creating Docker network $NETWORK_NAME..."
docker network create $NETWORK_NAME

# Check if the Zookeeper container exists; if so, remove it
if docker ps -a --format '{{.Names}}' | grep -q "^${ZOOKEEPER_CONTAINER_NAME}$"; then
  echo "Zookeeper container already exists. Removing it..."
  docker rm -f $ZOOKEEPER_CONTAINER_NAME
fi

# Start Zookeeper container
echo "Starting Zookeeper container..."
docker run -d --name $ZOOKEEPER_CONTAINER_NAME \
  --network $NETWORK_NAME \
  -p $ZOOKEEPER_PORT:2181 \
  -e ALLOW_ANONYMOUS_LOGIN=yes \
  bitnami/zookeeper:latest

# Check if the Kafka container exists; if so, remove it
if docker ps -a --format '{{.Names}}' | grep -q "^${KAFKA_CONTAINER_NAME}$"; then
  echo "Kafka container already exists. Removing it..."
  docker rm -f $KAFKA_CONTAINER_NAME
fi

# Start Kafka container
echo "Starting Kafka container..."
docker run -d --name $KAFKA_CONTAINER_NAME \
  --network $NETWORK_NAME \
  -p $KAFKA_PORT:9092 \
  -e KAFKA_BROKER_ID=1 \
  -e KAFKA_ZOOKEEPER_CONNECT=$ZOOKEEPER_CONTAINER_NAME:$ZOOKEEPER_PORT \
  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:$KAFKA_PORT \
  -e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT \
  bitnami/kafka:latest

echo "Kafka and Zookeeper are starting up..."
echo "Kafka is available at localhost:$KAFKA_PORT"

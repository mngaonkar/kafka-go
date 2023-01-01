package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("no hostname, error = ", err)
		os.Exit(1)
	}
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "147.182.230.42:9092",
		"client.id":         hostname,
		"acks":              "all",
	})

	if err != nil {
		log.Fatal("error initializting kafka producer, error = ", err)
		os.Exit(1)
	} else {
		log.Println("kafka producer initialized ", p)
	}

	topic := "test-messages"
	message := "first message from go producer"
	delivery_channel := make(chan kafka.Event, 10000)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(message),
	}, delivery_channel)
	if err != nil {
		log.Fatal("error sending message, error = ", err)
	} else {
		log.Println("message sent -> ", message)
	}
}

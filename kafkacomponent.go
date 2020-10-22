package main

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func produceToKafka(messages string) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "test-topic-a"
	welcomeMesagge := []string{messages, messages}

	allMessages := append([]string{messages}, welcomeMesagge...)
	for {
		for _, word := range allMessages {

			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(word),
			}, nil)
		}
		// Wait for message deliveries before shutting down
		p.Flush(15 * 1000)
		time.Sleep(time.Duration(rand.Int31n(3)) * time.Second)
	}

}

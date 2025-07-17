package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// Kafka broker addresses and topic name
	brokers := []string{"localhost:9092"} // Replace with your Kafka broker address
	topic := "example-topic"

	// Kafka configuration
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll              // Wait for all replicas to acknowledge
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner // select partitioner
	config.Producer.Return.Successes = true                       // Ensure success message delivery
	config.Version = sarama.V2_8_0_0                              // Specify the Kafka version to ensure compatibility

	// Create a Kafka producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Failed to close Kafka producer: %v", err)
		}
	}()

	// Use a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Start producing messages
	produceMessages(ctx, producer, topic)
}

func produceMessages(ctx context.Context, producer sarama.SyncProducer, topic string) {
	// Initialize the message index
	messageIndex := 0

	// Create a ticker to simulate 1-second intervals
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Produce messages until the context timeout occurs
	for {
		select {
		case <-ctx.Done():
			// Context timeout reached, end the message production
			fmt.Println("Stopping producer after timeout...")
			return
		case <-ticker.C:
			// Produce a message every second
			message := &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(fmt.Sprintf("Message #%d", messageIndex)),
			}

			// Send the message to the Kafka topic
			partition, offset, err := producer.SendMessage(message)
			if err != nil {
				log.Printf("Failed to send message #%d: %v", messageIndex, err)
				continue
			}

			// Log success
			fmt.Printf("Successfully sent message #%d to partition %d at offset %d\n", messageIndex, partition, offset)

			// Increment message index
			messageIndex++
		}
	}
}

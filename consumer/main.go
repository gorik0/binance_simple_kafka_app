package main

import (
	"context"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL)

	ctx, cancel := context.WithCancel(context.Background())

	// go routine for getting signals asynchronously
	go func() {
		sig := <-signals
		fmt.Println("Got signal: ", sig)
		cancel()
	}()

	addr := "localhost:29092"
	topic := "coins"

	config := kafka.ReaderConfig{
		Brokers: []string{addr},

		Topic: topic,
	}

	r := kafka.NewReader(config)

	fmt.Println("Consumer configuration: ", config)

	defer func() {
		err := r.Close()
		if err != nil {
			fmt.Println("Error closing consumer: ", err)
			return
		}
		fmt.Println("Consumer closed")
	}()

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Error reading message: ", err)
			break
		}
		fmt.Printf("Received message from %s-%d [%d]: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}

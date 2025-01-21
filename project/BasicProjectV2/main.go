package main

import "log"

func main() {
	app := InitializeApp()

	server := app.server
	sarama := app.saramaConsumer
	go func() {
		if err := sarama.ConsumeReadEvent(); err != nil {
			log.Fatalf("Failed to consume Kafka events: %v", err)
		}
	}()
	err := server.Run(":8088")
	if err != nil {
		return
	}

}

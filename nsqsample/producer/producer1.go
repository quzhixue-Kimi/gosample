package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("10.68.180.137:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i <= 100; i++ {
		messageBody := []byte("hello")
		topicName := "whatisup"

		err = producer.Publish(topicName, messageBody)
		if err != nil {
			log.Fatal(err)
		}
	}

	producer.Stop()
}

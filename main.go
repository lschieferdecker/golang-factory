package main

import (
	"fmt"
)

func main() {
	messageChannel := fakeMessageQueue()

	for _, msg := range messageChannel {
		// gets the consumer by calling the factory
		consumer, err := getConsumer(msg.Type)
		if err != nil {
			fmt.Printf("Failed to consume message: %v. Error: %v\n", msg, err.Error())
			continue
		}

		// consume the message
		consumer.consume(msg)
	}
}

// fakeMessageQueue returns an slice of message that represents our queue
// that will be consumed.
func fakeMessageQueue() []message {
	return []message{
		message{Type: "REGULAR", Text: "Golang"}, // will invoke regularConsumer
		message{Type: "REVERSE", Text: "Golang"}, // will invoke reverseConsumer
		message{Type: "ANY", Text: "Golang"},     // will throw error because there is no consumer
	}
}

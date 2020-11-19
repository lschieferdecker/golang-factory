package main

import (
	"fmt"
)

// getConsumer is a consumer factory
func getConsumer(messageType string) (iConsumer, error) {
	switch messageType {
	case "REGULAR":
		return newRegularConsumer(), nil
	case "REVERSE":
		return newReverseConsumer(), nil
	default:
		return nil, fmt.Errorf("Consumer not implemented")
	}
}

package main

import (
	"fmt"
)

// regularConsumer is a concrete product
type regularConsumer struct{}

func (rc *regularConsumer) consume(msg message) {
	fmt.Printf("%v\n", msg.Text)
}

// newRegularConsumer creates a new regularConsumer struct
func newRegularConsumer() iConsumer {
	return &regularConsumer{}
}

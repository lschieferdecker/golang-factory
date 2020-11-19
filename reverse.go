package main

import (
	"fmt"
)

// reverseConsumer is a concrete product
type reverseConsumer struct{}

func (rc *reverseConsumer) consume(msg message) {
	reverse := ""
	for _, c := range msg.Text {
		reverse = string(c) + reverse
	}
	fmt.Printf("%v\n", reverse)
}

// newReverseConsumer creates a reverseConsumer struct
func newReverseConsumer() iConsumer {
	return &reverseConsumer{}
}

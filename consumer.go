package main

// iConsumer is the product interface
type iConsumer interface {
	consume(msg message)
}

// Message is the event message that our consumer will consume
type message struct {
	Type string
	Text string
}

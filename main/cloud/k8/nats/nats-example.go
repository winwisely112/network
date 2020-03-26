package main

import (
	"log"

	"github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats-streaming"
)

func main() {
	nc, err := nats.Connect("nats://example-nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	sc, err := stan.Connect("example-stan", "client-123", stan.NatsConn(nc))
	if err != nil {
		log.Fatal(err)
	}
	sc.Publish("hello", []byte("one"))
	sc.Publish("hello", []byte("two"))
	sc.Publish("hello", []byte("three"))

	sc.Subscribe("hello", func(m *stan.Msg) {
		log.Printf("[Received] %+v", m)
	}, stan.DeliverAllAvailable())

	select {}
}

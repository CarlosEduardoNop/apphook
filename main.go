package main

import (
	"fmt"
	"time"

	"github.com/CarlosEduardoNop/apphook/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

func main() {
	for i := 0; i < 1000; i++ {
		Publish()
		Publish()
		Publish()
		Publish()
		Publish()
		Publish()
		Publish()
		Publish()
		Publish()
	}
}

func Publish() {
	time.Sleep(1 * time.Second)
	ch, _ := rabbitmq.OpenChannel()

	ch.Publish(
		"",
		"send-apphook",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(`{"url": "https://webhook.site/5f190f9f-5a50-4b82-af70-93413a32043d", "event": "orderCreated", "payload": {"name": "Carlos", "price": 1000}, "delay": 0}`),
		},
	)

	fmt.Println("Message sent")
}

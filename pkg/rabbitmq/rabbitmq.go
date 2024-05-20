package rabbitmq

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type CallbackFunc func(msg []byte)

func OpenChannel() (*amqp.Channel, error) {
	godotenv.Load()

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_PREFIX") + os.Getenv("RABBITMQ_USER") + ":" + os.Getenv("RABBITMQ_PASSWORD") + os.Getenv("RABBITMQ_HOST") + os.Getenv("RABBITMQ_PORT"))

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	return ch, nil
}

func Consume(ch *amqp.Channel, callback CallbackFunc, queueName string) {
	msgs, _ := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			callback(d.Body)
		}
	}()

	fmt.Println("Worker Started")

	<-forever
}

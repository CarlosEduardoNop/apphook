package main

import (
	"encoding/json"
	"fmt"

	"github.com/CarlosEduardoNop/apphook/pkg/rabbitmq"
)

func main() {
	for i := 0; i < 10; i++ {
		go Worker()
	}
	Worker()
}

func Worker() {
	ch, _ := rabbitmq.OpenChannel()

	rabbitmq.Consume(ch, func(msg []byte) {
		var input map[string]interface{}
		err := json.Unmarshal(msg, &input)

		if err != nil {
			fmt.Println(err)
			return
		}

		body, err := json.Marshal(input["payload"])

		if err != nil {
			fmt.Println(err)
			return
		}

		url := input["url"].(string)

		go MakeRequest(url, body, input)

		fmt.Println("Message received")
	}, "send-apphook")
}

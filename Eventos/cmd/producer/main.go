package main

import (
	"github.com/Jorge79/estudos-go/Eventos/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	rabbitmq.Publish(ch, "Hello World", "amq.direct")
}

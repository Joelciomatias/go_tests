package main

import (
	"github.com/streadway/amqp"
)

func main() {
	print("Helllo from publisher")
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672")
	ch, _ := conn.Channel()

	q, _ := ch.QueueDeclare(
		"first.queue",
		true,
		false,
		false,
		false,
		nil)

	// ch.QueueBind(
	// 	q.Name,
	// 	"",
	// 	"amq.fanout",
	// 	false,
	// 	nil)

	msg := amqp.Publishing{
		Body: []byte("my first message"),
	}

	for {
		ch.Publish(
			"",
			q.Name,
			false,
			false,
			msg)
	}

}

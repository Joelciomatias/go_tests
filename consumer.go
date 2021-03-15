package main

import (
	"github.com/streadway/amqp"
)

func main() {
	print("Helllo")
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672")
	ch, _ := conn.Channel()

	q, _ := ch.QueueDeclare(
		"first.queue",
		true,
		false,
		false,
		false,
		nil)
	// fmt.Printf(q)
	// ch.QueueBind(
	// 	q.Name,
	// 	"",
	// 	"amq.fanout",
	// 	false,
	// 	nil)

	msgs, _ := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)

	for m := range msgs {
		println(string(m.Body))
	}

}

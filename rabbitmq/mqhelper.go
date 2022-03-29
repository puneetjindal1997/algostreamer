package rabbitmq

import (
	"encoding/json"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Send(queueName string, data map[string]interface{}) {
	log.Println("queueName, data", queueName, data)
	// conn, err := amqp.Dial("amqp://" + os.Getenv("MQUserName") + ":" + os.Getenv("MQPasword") + "@" + os.Getenv("MQHOST") + ":" + os.Getenv("MQPORT") + "/")
	conn, err := amqp.Dial("amqp://" + os.Getenv("MQUserName") + ":" + os.Getenv("MQPasword") + "@" + os.Getenv("MQHOST") + ":" + os.Getenv("MQPORT") + "/")
	// conn, err := amqp.Dial("amqp://guest:guest@rabbitmq_c")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")
	bodyDataByte, _ := json.Marshal(data)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bodyDataByte,
		})
	failOnError(err, "Failed to declare a queue")
}

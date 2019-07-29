package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func publish() {

}

func getConnection() *amqp.Connection {
	amqpHost := getEnv("AMQP_HOST", "localhost")
	amqpPort := getEnv("AMQP_PORT", "5672")
	amqpUser := getEnv("AMQP_USER", "guest")
	amqpPwd := getEnv("AMQP_PWD", "guest")

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", amqpUser, amqpPwd, amqpHost, amqpPort))
	if err != nil {
		fmt.Println("failed init connection to amqp")
	}
	return conn
}

package main

import (
	"fmt"
	"time"

	"github.com/kong11213613/haremicro/broker"
	"github.com/kong11213613/haremicro/broker/asynq"
	"github.com/kong11213613/haremicro/logger"
)

var (
	topic = "haremicro.topic"
)

func pub() {
	tick := time.NewTicker(5 * time.Second)
	i := 0
	for _ = range tick.C {
		queue := "default"
		if i%4 == 0 {
			queue = "critical"
		}
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := broker.Publish(topic, msg, asynq.Queue(queue)); err != nil {
			logger.Errorf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
	}
}

func main() {
	broker.DefaultBroker = asynq.NewBroker(
		asynq.DB(1),
		asynq.Service("test"),
	)

	if err := broker.Init(); err != nil {
		logger.Fatalf("Broker Init error: %v", err)
	}

	if err := broker.Connect(); err != nil {
		logger.Fatalf("Broker Connect error: %v", err)
	}

	pub()
}

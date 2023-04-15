package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"pro/middleware/log"
	"time"
)

func Queue1Consume() {
	go Subscribe(Queue1, func(msgs <-chan amqp.Delivery) {
		for msg := range msgs {
			go func(msg amqp.Delivery) {
				log.Info("收到消息时间: ", time.Now().Unix())
				fmt.Printf("%s 收到消息：%v\n", Queue1, string(msg.Body))
				msg.Ack(false)
				//重新入列
				//msg.Reject(true)
			}(msg)
		}
	})
}

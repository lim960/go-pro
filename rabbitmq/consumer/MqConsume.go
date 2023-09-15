package consume

import (
	"fmt"
	"github.com/streadway/amqp"
	"pro/rabbitmq/config"
)

func Subscribe(queue string, callback func(<-chan amqp.Delivery)) {
	msgs, err := mqConfig.GetChannel().Consume(queue, queue, false, false, false, false, nil)

	if err != nil {
		panic(err)
	}
	callback(msgs)
}

func InItConsume() {

	Queue1Consume()
	println("mq消费者注册完成")
}

func Queue1Consume() {
	go Subscribe(mqConfig.Queue1, func(msgs <-chan amqp.Delivery) {
		for msg := range msgs {
			go func(msg amqp.Delivery) {
				//log.Info("收到消息时间: ", time.Now().Unix())
				data := string(msg.Body)
				fmt.Printf("%s 收到消息：%v\n", mqConfig.Queue1, data)
				defer msg.Ack(false)

				//重新入列
				//msg.Reject(true)
			}(msg)
		}
	})
}

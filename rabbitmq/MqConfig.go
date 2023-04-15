package rabbitmq

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
	"net/url"
	logs "pro/middleware/log"
	"time"
)

var conn *amqp.Connection

var ch *amqp.Channel

var err error

const (
	DelayExchange = "delayExchange"
	Queue1        = "queue1"
)

func InitMq() {
	host := viper.GetString("rabbitmq.host")
	port := viper.GetString("rabbitmq.port")
	username := viper.GetString("rabbitmq.username")
	password := viper.GetString("rabbitmq.password")
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		username,
		url.QueryEscape(password),
		host,
		port)
	conn, err = amqp.Dial(dsn)
	//ch = conn
	if err != nil {
		panic(err)
	}

	ch, err = conn.Channel()
	if err != nil {
		panic(err)
	}
	//申明交换机
	err := ch.ExchangeDeclare(DelayExchange, "x-delayed-message",
		//交换机持久化
		false,
		false,
		false,
		false,
		map[string]interface{}{"x-delayed-type": "direct"})
	if err != nil {
		log.Fatal(err)
	}
	//初始化队列
	InitDelayMQ(Queue1)
	// 限制未ack的最多有10个,必须设置为手动ack才有效
	ch.Qos(10, 0, false)

	//启动消费者
	Queue1Consume()
}

func InitDelayMQ(queue string) {
	// 声明 queue
	_, err = ch.QueueDeclare(queue,
		//队列持久化
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}
	// 将 queue 与 exchange绑定
	err = ch.QueueBind(queue, queue, DelayExchange, false, nil)
	if err != nil {
		panic(err)
	}
}

// SendDelayMessage 发送延时消息
// delay: 延迟时间 单位秒
func SendDelayMessage(queue string, message string, delay int) {
	logs.Info("发送时间: ", time.Now().Unix())
	err := ch.Publish(DelayExchange, queue, true, false, amqp.Publishing{
		Headers: map[string]interface{}{"x-delay": delay * 1000},
		Body:    []byte(fmt.Sprintf("%v", message)),
	})
	if err != nil {
		panic(err)
	}
}

// SendPersistentDelayMessage 发送持久化延时消息
// delay: 延迟时间 单位秒
func SendPersistentDelayMessage(queue string, message string, delay int) {
	logs.Info("发送时间: ", time.Now().Unix())
	err := ch.Publish(DelayExchange, queue, true, false, amqp.Publishing{
		Headers:      map[string]interface{}{"x-delay": delay * 1000},
		Body:         []byte(fmt.Sprintf("%v", message)),
		DeliveryMode: 2,
	})
	if err != nil {
		panic(err)
	}
}

func Subscribe(queue string, callback func(<-chan amqp.Delivery)) {
	msgs, err := ch.Consume(queue, queue, false, false, false, false, nil)

	if err != nil {
		panic(err)
	}
	callback(msgs)
}

//func InitMQ(ch *amqp.Channel, queue, key, exchange string) {
//	// 声明 exchange
//	err := ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)
//	if err != nil {
//		panic(err)
//	}
//	// 声明 queue
//	_, err = ch.QueueDeclare(queue, false, false, false, false, nil)
//	if err != nil {
//		panic(err)
//	}
//	// 将 queue 与 exchange 和 key 绑定
//	err = ch.QueueBind(queue, key, exchange, false, nil)
//	if err != nil {
//		panic(err)
//	}
//
//}
//
//func sendMessage(ch *amqp.Channel, exchange string, key string, message string) {
//	err := ch.Publish(exchange, key, true, false, amqp.Publishing{
//		Body: []byte(fmt.Sprintf("%v", message)),
//	})
//	if err != nil {
//		panic(err)
//	}
//
//}
//
//func SetConfirm(ch *amqp.Channel, notifyConfirm chan amqp.Confirmation) {
//	err := ch.Confirm(false)
//	if err != nil {
//		log.Println(err)
//	}
//	notifyConfirm = ch.NotifyPublish(make(chan amqp.Confirmation))
//}
//
//func ListenConfirm(notifyConfirm chan amqp.Confirmation) {
//	for ret := range notifyConfirm {
//		if ret.Ack {
//			fmt.Println("消息发送成功")
//		} else {
//			fmt.Println("消息发送失败")
//		}
//	}
//}
//
//func NotifyReturn(notifyReturn chan amqp.Return, channel *amqp.Channel) {
//	notifyReturn = channel.NotifyReturn(make(chan amqp.Return))
//}
//func ListReturn(notifyReturn chan amqp.Return) {
//	ret := <-notifyReturn
//	if string(ret.Body) != "" {
//		fmt.Println("消息没有投递到队列:", string(ret.Body))
//		panic("skfh")
//	}
//}

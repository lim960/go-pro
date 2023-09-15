package mqConfig

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
	Queue1        = "mQueue1"
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
	println("mq注册完成")
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

func GetChannel() *amqp.Channel {
	return ch
}

// SendDelayMessage 发送延时消息
// delay: 延迟时间 单位秒
func SendDelayMessage(queue string, message string, delay int) {
	//logs.Info("发送时间: ", time.Now().Unix())
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

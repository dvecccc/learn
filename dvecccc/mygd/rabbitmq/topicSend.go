package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var user = "admin"
var pwd = "wang8866."
var host = "47.109.33.47"
var port = "5672"
var MqURL = "amqp://" + user + ":" + pwd + "@" + host + ":" + port + "/my_vhost"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	key       string
	MqURL     string
}

func (r *RabbitMQ) failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (r *RabbitMQ) CloseConn() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, key: key, MqURL: MqURL}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqURL)
	rabbitmq.failOnError(err, "创建连接错误")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnError(err, "获取channel失败")
	return rabbitmq
}

func NewRabbitMQTopic(exchangeName, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqURL)
	rabbitmq.failOnError(err, "连接rabbitmq失败")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnError(err, "failed to open a channel")
	return rabbitmq
}

func (r *RabbitMQ) PublishTopic(message string) {
	//尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnError(err, "failed to declare an exchange")
	//发送消息
	err = r.channel.Publish(
		r.Exchange,
		r.key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func (r *RabbitMQ) ReceiveTopic() {
	//尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnError(err, "failed to declare an exchange")
	//尝试创建队列，队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生成队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnError(err, "failed to declare an exchange")
	//绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		//需要绑定key
		r.key,
		r.Exchange,
		false,
		nil,
	)
	//消费消息
	message, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range message {
			log.Printf("received a message : %s", d.Body)
		}
	}()

	fmt.Println("[*] Waiting for message, To exit press CTRL+C")
	<-forever
}

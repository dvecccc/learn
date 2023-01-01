package rabbitmq

import (
	"fmt"
	"strconv"
	"time"
)

func Publish() {
	t1 := NewRabbitMQTopic("chain_one", "microA.1")
	//t2 := NewRabbitMQTopic("chain_two", "microB")
	//t3 := NewRabbitMQTopic("chain_three", "microA")
	//t4 := NewRabbitMQTopic("chain_four", "microB")
	t5 := NewRabbitMQTopic("chain_one", "microA.2")
	r1 := NewRabbitMQTopic("chain_one", "#")
	defer t1.CloseConn()
	defer t5.CloseConn()
	defer r1.CloseConn()
	go r1.ReceiveTopic()
	time.Sleep(time.Second)
	for i := 1; i < 10; i++ {
		t1.PublishTopic(fmt.Sprintf("1 chan_one, keyA, 第%s条消息", strconv.Itoa(i)))
		//fmt.Println("www")
		time.Sleep(time.Second)
		//t2.PublishTopic(fmt.Sprintf("2 chan_two, keyB, 第%s条消息", strconv.Itoa(i)))
		//time.Sleep(time.Second)
		//t3.PublishTopic(fmt.Sprintf("3 chan_three, keyA, 第%s条消息", strconv.Itoa(i)))
		//time.Sleep(time.Second)
		//t4.PublishTopic(fmt.Sprintf("4 chan_four, keyB, 第%s条消息", strconv.Itoa(i)))
		//time.Sleep(time.Second)
		t5.PublishTopic(fmt.Sprintf("5 chan_one, keyA, 第%s条消息", strconv.Itoa(i)))
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second * 10)
}

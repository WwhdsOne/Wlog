package mq

type Producer interface {

	// 初始化生产者
	initProducer()

	// 发送消息
	sendMessage(messageValue string)
}

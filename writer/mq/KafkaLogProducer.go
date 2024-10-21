package mq

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

type KafkaLogProducer struct {
	topic     string `yaml:"topic"`
	partition int    `yaml:"partition"`
	host      string `yaml:"host"`
	port      int    `yaml:"port"`
	producer  sarama.SyncProducer
}

// 初始化 Kafka 生产者
func (k *KafkaLogProducer) initProducer() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll

	// 使用 host 和 port 字段拼接 Kafka broker 地址
	brokerAddress := fmt.Sprintf("%s:%d", k.host, k.port)
	producer, err := sarama.NewSyncProducer([]string{brokerAddress}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
	k.producer = producer
}

// 发送消息
func (k *KafkaLogProducer) sendMessage(messageValue string) {
	message := &sarama.ProducerMessage{
		Topic:     k.topic,            // 使用结构体中的 topic
		Partition: int32(k.partition), // 使用结构体中的 partition
		Value:     sarama.StringEncoder(messageValue),
	}

	partition, offset, err := k.producer.SendMessage(message)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	} else {
		fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	}
}

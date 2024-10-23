package mq

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/WwhdsOne/Wlog/WLog"
	"log"
)

type KafkaLogProducer struct {
	Topic     string `yaml:"topic"`
	Partition int    `yaml:"partition"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	producer  sarama.SyncProducer
}

// InitWriter InitProducer 初始化 Kafka 生产者
func (k *KafkaLogProducer) InitWriter() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll

	// 使用 host 和 port 字段拼接 Kafka broker 地址
	brokerAddress := fmt.Sprintf("%s:%d", k.Host, k.Port)
	producer, err := sarama.NewSyncProducer([]string{brokerAddress}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
	k.producer = producer
}

func (k *KafkaLogProducer) Write(p []byte) (n int, err error) {
	message := &sarama.ProducerMessage{
		Topic:     k.Topic,                 // 使用结构体中的 topic
		Partition: int32(k.Partition),      // 使用结构体中的 partition
		Value:     sarama.StringEncoder(p), // 使用传入的 []byte
	}
	_, _, err = k.producer.SendMessage(message)
	if err != nil {
		WLog.Error(err.Error())
		return 0, err
	}
	return len(p), nil
}

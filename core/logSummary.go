package core

import (
	"zapLog/writer/file"
	"zapLog/writer/mq"
)

type summaryConfig struct {

	// 本地文件日志写入
	writers []file.Writer `yaml:"writers"`

	// 向各类消息队列发送日志信息
	producers []mq.Producer `yaml:"producers"`
}

package WLog

import (
	"github.com/WwhdsOne/Wlog/writer"
)

type LogSummary struct {
	LocalFileWriter writer.LogWriter `yaml:"localFileWriter"` // 本地文件日志写入
	KafkaWriter     writer.LogWriter `yaml:"kafkaWriter"`     // 向各类消息队列发送日志信息
	LogFormatConfig *LogFormatConfig `yaml:"logFormatConfig"` // 日志格式配置
}

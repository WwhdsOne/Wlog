package WLog

import (
	"github.com/WwhdsOne/Wlog/writer"
)

type LogSummary struct {
	LocalFileWriter writer.LogWriter // 本地文件日志写入
	KafkaWriter     writer.LogWriter // 向各类消息队列发送日志信息
	SysLogWriter    writer.LogWriter // syslog形式的日志
	LogFormatConfig *LogFormatConfig // 日志格式配置
}

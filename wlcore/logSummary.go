package wlcore

import (
	"dev.aminer.cn/codegeex-enterprise/wlog/writer"
	"go.uber.org/zap/zapcore"
	"os"
)

type LogSummary struct {
	LocalFileWriter writer.LogWriter // 本地文件日志写入
	KafkaWriter     writer.LogWriter // 向各类消息队列发送日志信息
	SysLogWriter    writer.LogWriter // syslog形式的日志
	LogFormatConfig *LogFormatConfig // 日志格式配置
	Rfc5424Config   *Rfc5424Config   // Rfc5424的日志配置
}

// BuildWriters 新增符合zap要求的Writer
func (ls *LogSummary) BuildWriters() []zapcore.WriteSyncer {
	// 创建多个输出目标
	var writers = make([]zapcore.WriteSyncer, 0)

	// 默认输出到控制台
	writers = append(writers, zapcore.AddSync(os.Stdout))

	// 如果配置了本地文件输出，则追加到输出目标
	if lfs := ls.LocalFileWriter; lfs != nil {
		lfs.InitWriter()
		writers = append(writers, zapcore.AddSync(lfs))
	}

	// 如果配置了kafka输出，则追加到输出目标
	if kw := ls.KafkaWriter; kw != nil {
		kw.InitWriter()
		writers = append(writers, zapcore.AddSync(kw))
	}

	// 如果配置了syslog输出，则追加到输出目标
	if sl := ls.SysLogWriter; sl != nil {
		sl.InitWriter()
		writers = append(writers, zapcore.AddSync(sl))
	}
	return writers
}

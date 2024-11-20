package writer

import (
	"os"

	"go.uber.org/zap/zapcore"
)

var DefaultWLogWriter = &WLogWriters{}

type WLogWriters struct {
	LocalFileWriter logWriter // 本地文件日志写入
	KafkaWriter     logWriter // 向各类消息队列发送日志信息
	SysLogWriter    logWriter // syslog形式的日志
}

// BuildWriters 新增符合zap要求的Writer
func (ls *WLogWriters) BuildWriters() []zapcore.WriteSyncer {
	// 创建多个输出目标
	var writers = make([]zapcore.WriteSyncer, 0)

	// 默认输出到控制台
	writers = append(writers, zapcore.AddSync(os.Stdout))

	// 如果配置了本地文件输出，则追加到输出目标
	if lfw := ls.LocalFileWriter; lfw != nil {
		lfw.InitWriter()
		writers = append(writers, zapcore.AddSync(lfw))
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

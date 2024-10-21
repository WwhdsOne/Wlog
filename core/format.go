package core

import "go.uber.org/zap/zapcore"

// LogFormatConfig 用于指定日志的各项格式化设置
type LogFormatConfig struct {
	// 日志级别
	Level zapcore.Level
	// 日志前缀
	Prefix string
	// 是否采用json格式
	IsJson bool
	// 编码格式
	EncoderLevel string
	// 打印堆栈信息的等级
	StacktraceLevel zapcore.LevelEnabler
}

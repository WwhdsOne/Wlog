package wlcore

import (
	"go.uber.org/zap/zapcore"
	"os"
)

// LogFormatConfig 用于指定日志的各项格式化设置
type LogFormatConfig struct {
	Level           zapcore.Level        // 日志级别
	Prefix          string               // 日志前缀
	IsJson          bool                 // 是否采用json格式
	EncoderLevel    string               // 编码格式
	StacktraceLevel zapcore.LevelEnabler // 打印堆栈信息的等级
}

// NewLogFormatConfig 初始化默认值
func NewLogFormatConfig() *LogFormatConfig {
	return &LogFormatConfig{
		Level:           zapcore.DebugLevel,
		Prefix:          "[" + os.Args[0] + "]",
		IsJson:          true,
		EncoderLevel:    "LowercaseLevelEncoder",
		StacktraceLevel: zapcore.ErrorLevel,
	}
}

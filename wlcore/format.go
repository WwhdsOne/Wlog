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
		Prefix:          os.Args[0],
		IsJson:          true,
		EncoderLevel:    "LowercaseLevelEncoder",
		StacktraceLevel: zapcore.ErrorLevel,
	}
}

type Rfc5424Config struct {
	Hostname string
	AppName  string
}

// FillEmptyLogFormat  设置默认日志格式
func (lfc *LogFormatConfig) FillEmptyLogFormat() {

	// 前缀为空则使用程序名
	if lfc.Prefix == "" {
		lfc.Prefix = os.Args[0]
	}

	// 编码等级为空则使用小写无色编码
	if lfc.EncoderLevel == "" {
		lfc.EncoderLevel = "LowercaseLevelEncoder"
	}

	// 堆栈跟踪等级为错误等级
	if lfc.StacktraceLevel == nil {
		lfc.StacktraceLevel = zapcore.ErrorLevel
	}
}

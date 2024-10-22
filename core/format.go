package core

import "go.uber.org/zap/zapcore"

// LogFormatConfig 用于指定日志的各项格式化设置
type LogFormatConfig struct {
	Level           zapcore.Level        `yaml:"level"`           // 日志级别
	Prefix          string               `yaml:"prefix"`          // 日志前缀
	IsJson          bool                 `json:"isJson"`          // 是否采用json格式
	EncoderLevel    string               `yaml:"encoderLevel"`    // 编码格式
	StacktraceLevel zapcore.LevelEnabler `yaml:"stacktraceLevel"` // 打印堆栈信息的等级
}

// NewLogFormatConfig 初始化默认值
func NewLogFormatConfig() *LogFormatConfig {
	return &LogFormatConfig{
		Level:           zapcore.DebugLevel,
		Prefix:          "[ZAP-JSON]",
		IsJson:          true,
		EncoderLevel:    "LowercaseLevelEncoder",
		StacktraceLevel: zapcore.ErrorLevel,
	}
}

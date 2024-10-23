package WLog

import (
	"go.uber.org/zap/zapcore"
	"time"
)

func Encoder(encodeLevel string, isJson bool) zapcore.Encoder {
	// 创建一个 zapcore.EncoderConfig 配置对象
	config := zapcore.EncoderConfig{
		TimeKey:        "time",                         // 时间字段的键名
		NameKey:        "name",                         // 日志记录器名称的键名
		LevelKey:       "level",                        // 日志级别的键名
		CallerKey:      "caller",                       // 调用者的键名
		MessageKey:     "msg",                          // 日志消息的键名
		StacktraceKey:  "stacktrace",                   // 堆栈跟踪的键名，从配置中获取
		LineEnding:     zapcore.DefaultLineEnding,      // 行尾字符，使用默认值
		EncodeLevel:    LevelEncoder(encodeLevel),      // 大写编码器带颜色
		EncodeCaller:   zapcore.FullCallerEncoder,      // 调用者编码器，使用完整路径
		EncodeDuration: zapcore.SecondsDurationEncoder, // 持续时间编码器，使用秒数
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) { // 时间编码器
			encoder.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
	}
	// 根据配置中的格式选择编码器
	if isJson == true {
		// 如果格式为 "json"，则返回 JSON 编码器
		return zapcore.NewJSONEncoder(config)
	}
	// 否则返回控制台编码器
	return zapcore.NewConsoleEncoder(config)
}

// LevelEncoder 根据 EncodeLevel 返回 zapcore.LevelEncoder
func LevelEncoder(EncodeLevel string) zapcore.LevelEncoder {
	switch {
	case EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

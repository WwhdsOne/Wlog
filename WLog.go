package WLog

import (
	"fmt"
	"github.com/WwhdsOne/Wlog/wlcore"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LowercaseLevelEncoder      = "LowercaseLevelEncoder"      // 小写编码器(默认)
	LowercaseColorLevelEncoder = "LowercaseColorLevelEncoder" // 小写编码器带颜色
	CapitalLevelEncoder        = "CapitalLevelEncoder"        // 大写编码器
	CapitalColorLevelEncoder   = "CapitalColorLevelEncoder"   // 大写编码器带颜色
)

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
)

type Logger struct {
	l       *zap.Logger           // 日志实体
	al      *zap.AtomicLevel      // 日志等级
	prefix  string                // 日志前缀
	rfc5424 *wlcore.Rfc5424Config // RFC5424配置
}

func Build(ls *wlcore.LogSummary) *Logger {

	// 获取日志设置
	lfc := ls.LogFormatConfig
	if lfc == nil {
		lfc = wlcore.NewLogFormatConfig()
	}
	lfc.FillEmptyLogFormat()

	// 初始化日志等级，有对应的动态调整接口
	al := zap.NewAtomicLevelAt(lfc.Level)

	// 初始化日志编码格式
	var encoder zapcore.Encoder
	if ls.Rfc5424Config == nil {
		encoder = wlcore.Encoder(lfc.EncoderLevel, lfc.IsJson, false)
	} else {
		encoder = wlcore.Encoder(lfc.EncoderLevel, lfc.IsJson, true)
	}

	// 初始化日志输出
	writers := ls.BuildWriters()

	// 创建日志输出对象
	Wzapcore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writers...), al)

	return &Logger{
		l:       zap.New(Wzapcore, zap.AddStacktrace(lfc.StacktraceLevel)),
		al:      &al,
		prefix:  "[" + lfc.Prefix + "]",
		rfc5424: ls.Rfc5424Config,
	}
}

func (l *Logger) formatMessage(msg string, loptions *wlcore.Loptions) string {
	if loptions.Package != "" {
		msg = fmt.Sprintf("package = %s | %s", loptions.Package, msg)
	}
	msg = l.prefix + " | " + msg
	if len(loptions.Option) != 0 {
		msg = fmt.Sprintf(msg, loptions.Option...)
	}
	if loptions.Rfc5424Message != nil {
		msg = loptions.Rfc5424Message.FormatRfc5424Message(l.rfc5424.Hostname, l.rfc5424.AppName, msg)
	}
	return msg
}

func (l *Logger) Debug(msg string, loptions *wlcore.Loptions) {
	l.l.Debug(l.formatMessage(msg, loptions))
}

func (l *Logger) Info(msg string, loptions *wlcore.Loptions) {
	l.l.Info(l.formatMessage(msg, loptions))
}

func (l *Logger) Warn(msg string, loptions *wlcore.Loptions) {
	l.l.Warn(l.formatMessage(msg, loptions))
}

func (l *Logger) Error(msg string, loptions *wlcore.Loptions) {
	l.l.Error(l.formatMessage(msg, loptions))
}

func (l *Logger) Panic(msg string, loptions *wlcore.Loptions) {
	l.l.Panic(l.formatMessage(msg, loptions))
}

func (l *Logger) Fatal(msg string, loptions *wlcore.Loptions) {
	l.l.Fatal(l.formatMessage(msg, loptions))
}

// Sync 同步日志
func (l *Logger) Sync() error {
	return l.l.Sync()
}

var std = Build(&wlcore.LogSummary{})

func Default() *Logger         { return std }
func ReplaceDefault(l *Logger) { std = l }

func Debug(msg string) { std.Debug(msg, &wlcore.Loptions{}) }
func Info(msg string)  { std.Info(msg, &wlcore.Loptions{}) }
func Warn(msg string)  { std.Warn(msg, &wlcore.Loptions{}) }
func Error(msg string) { std.Error(msg, &wlcore.Loptions{}) }
func Panic(msg string) { std.Panic(msg, &wlcore.Loptions{}) }
func Fatal(msg string) { std.Fatal(msg, &wlcore.Loptions{}) }

func Sync() error { return std.Sync() }

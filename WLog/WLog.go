package WLog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
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
	l      *zap.Logger
	al     *zap.AtomicLevel
	prefix string // 日志前缀
}

func Build(ls *LogSummary) *Logger {

	// 获取日志
	lfc := fillEmptyLogFormat(ls.LogFormatConfig)
	// 初始化日志等级，有对应的动态调整接口

	al := zap.NewAtomicLevelAt(lfc.Level)

	// 初始化日志编码格式
	encoder := Encoder(lfc.EncoderLevel, lfc.IsJson)

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

	// 创建日志输出对象
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writers...), al)

	return &Logger{
		l:      zap.New(core, zap.AddStacktrace(lfc.StacktraceLevel)),
		al:     &al,
		prefix: "[" + lfc.Prefix + "]",
	}
}

// SetDefaultLogFormat 设置默认日志格式
func fillEmptyLogFormat(lfc *LogFormatConfig) *LogFormatConfig {
	if lfc == nil {
		return NewLogFormatConfig()
	}
	// 前缀为空则使用程序名
	if lfc.Prefix == "" {
		lfc.Prefix = os.Args[0]
	}
	// 编码等级为空则使用小写五色编码
	if lfc.EncoderLevel == "" {
		lfc.EncoderLevel = "LowercaseLevelEncoder"
	}
	// 堆栈跟踪等级为错误等级
	if lfc.StacktraceLevel == nil {
		lfc.StacktraceLevel = ErrorLevel
	}
	return lfc
}

// SetLevel 设置日志等级
func (l *Logger) SetLevel(level zapcore.Level) {
	if l.al != nil {
		l.al.SetLevel(level)
	}
}

func (l *Logger) formatMessage(msg string, loptions *Loptions) string {
	if loptions.Package != "" {
		msg = fmt.Sprintf("package = %s | %s", loptions.Package, msg)
	}
	msg = l.prefix + " | " + msg
	if len(loptions.Option) != 0 {
		msg = fmt.Sprintf(msg, loptions.Option...)
	}
	return msg
}

func (l *Logger) Debug(msg string, loptions *Loptions) {
	l.l.Debug(l.formatMessage(msg, loptions))
}

func (l *Logger) Info(msg string, loptions *Loptions) {
	l.l.Info(l.formatMessage(msg, loptions))
}

func (l *Logger) Warn(msg string, loptions *Loptions) {
	l.l.Warn(l.formatMessage(msg, loptions))
}

func (l *Logger) Error(msg string, loptions *Loptions) {
	l.l.Error(l.formatMessage(msg, loptions))
}

func (l *Logger) Panic(msg string, loptions *Loptions) {
	l.l.Panic(l.formatMessage(msg, loptions))
}

func (l *Logger) Fatal(msg string, loptions *Loptions) {
	l.l.Fatal(l.formatMessage(msg, loptions))
}

// Sync 同步日志
func (l *Logger) Sync() error {
	return l.l.Sync()
}

var std = Build(&LogSummary{})

func Default() *Logger         { return std }
func ReplaceDefault(l *Logger) { std = l }

func Debug(msg string) { std.Debug(msg, &Loptions{}) }
func Info(msg string)  { std.Info(msg, &Loptions{}) }
func Warn(msg string)  { std.Warn(msg, &Loptions{}) }
func Error(msg string) { std.Error(msg, &Loptions{}) }
func Panic(msg string) { std.Panic(msg, &Loptions{}) }
func Fatal(msg string) { std.Fatal(msg, &Loptions{}) }

func Sync() error { return std.Sync() }

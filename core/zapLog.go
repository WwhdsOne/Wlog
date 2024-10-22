package core

import (
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
	l  *zap.Logger
	al *zap.AtomicLevel
}

func Build(ls *LogSummary) *Logger {

	// 获取日志
	lfc := ls.LogFormatConfig
	if lfc == nil {
		lfc = NewLogFormatConfig()
	}
	// 初始化日志等级，有对应的动态调整接口

	al := zap.NewAtomicLevelAt(lfc.Level)

	// 初始化日志编码格式
	encoder := Encoder(lfc.Prefix, lfc.EncoderLevel, lfc.IsJson)

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
		l:  zap.New(core, zap.AddStacktrace(lfc.StacktraceLevel)),
		al: &al,
	}
}

func (l *Logger) SetLevel(level zapcore.Level) {
	if l.al != nil {
		l.al.SetLevel(level)
	}
}

type Field = zap.Field

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
	l.l.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.l.Fatal(msg, fields...)
}

// Sync 同步日志
func (l *Logger) Sync() error {
	return l.l.Sync()
}

var std = Build(&LogSummary{})

func Default() *Logger         { return std }
func ReplaceDefault(l *Logger) { std = l }

func Debug(msg string, fields ...Field) { std.Debug(msg, fields...) }
func Info(msg string, fields ...Field)  { std.Info(msg, fields...) }
func Warn(msg string, fields ...Field)  { std.Warn(msg, fields...) }
func Error(msg string, fields ...Field) { std.Error(msg, fields...) }
func Panic(msg string, fields ...Field) { std.Panic(msg, fields...) }
func Fatal(msg string, fields ...Field) { std.Fatal(msg, fields...) }

func Sync() error { return std.Sync() }

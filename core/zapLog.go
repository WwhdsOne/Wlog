package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"zapLog/config"
)

type Level = zapcore.Level

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

func Build(out []io.Writer, l config.LogFormatConfig) *Logger {

	// 初始化日志等级，有对应的动态调整接口
	al := zap.NewAtomicLevelAt(l.Level)

	// 初始化日志编码格式
	encoder := Encoder(l.Prefix, l.EncoderLevel, l.IsJson)

	// 创建多个输出目标
	var writers = make([]zapcore.WriteSyncer, 0, len(out))

	// 默认输出到控制台
	writers = append(writers, zapcore.AddSync(os.Stdout))

	// 添加自定义输出目标
	for _, w := range out {
		writers = append(writers, zapcore.AddSync(w))
	}

	// 创建日志输出对象
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writers...), al)

	return &Logger{
		l:  zap.New(core, zap.AddStacktrace(l.StacktraceLevel)),
		al: &al,
	}
}

func (l *Logger) SetLevel(level Level) {
	if l.al != nil {
		l.al.SetLevel(level)
	}
}

type Field = zap.Field

func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

var std = Build([]io.Writer{}, InfoLevel, "[Zap]", false)

func Default() *Logger         { return std }
func ReplaceDefault(l *Logger) { std = l }

func SetLevel() { std.SetLevel(level) }

func Debug(msg string, fields ...Field) { std.Debug(msg, fields...) }
func Info(msg string, fields ...Field)  { std.Info(msg, fields...) }
func Warn(msg string, fields ...Field)  { std.Warn(msg, fields...) }
func Error(msg string, fields ...Field) { std.Error(msg, fields...) }
func Panic(msg string, fields ...Field) { std.Panic(msg, fields...) }
func Fatal(msg string, fields ...Field) { std.Fatal(msg, fields...) }

func Sync() error { return std.Sync() }

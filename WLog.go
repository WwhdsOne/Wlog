package WLog

import (
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"dev.aminer.cn/codegeex-enterprise/cgxlog/cgx-wlog-go.git/opt"
	"dev.aminer.cn/codegeex-enterprise/cgxlog/cgx-wlog-go.git/writer"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

type Level int8

type Logger struct {
	l   *zap.Logger  // 日志实体
	Opt opt.Option   // 日志选项接口
	rw  sync.RWMutex // opt切换读写锁
}

func Build(level string, ls *writer.WLogWriters, option opt.Option) *Logger {

	level = strings.ToLower(level)
	// 初始化日志输出
	writers := ls.BuildWriters()

	// 创建日志输出对象

	lv := zap.NewAtomicLevelAt(zap.DebugLevel)
	switch level {
	case "debug":
		lv = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		lv = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		lv = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		lv = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	wZapCore := zapcore.NewCore(
		newEncoder(),                            // 编码器
		zapcore.NewMultiWriteSyncer(writers...), // 写入对象
		lv,
	)
	// 默认Rfc5424
	if option == nil {
		option = &opt.Rfc5424Opt{}
	}

	return &Logger{
		l:   zap.New(wZapCore, zap.AddStacktrace(zapcore.ErrorLevel)),
		Opt: option,
	}
}

func newEncoder() zapcore.Encoder {
	// 创建一个 zapcore.EncoderConfig 配置对象
	config := zapcore.EncoderConfig{
		NameKey:        "name",                         // 日志记录器名称的键名
		CallerKey:      "caller",                       // 调用者的键名
		MessageKey:     "msg",                          // 日志消息的键名
		StacktraceKey:  "stacktrace",                   // 堆栈跟踪的键名，从配置中获取
		LineEnding:     zapcore.DefaultLineEnding,      // 行尾字符，使用默认值
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 大写编码器带颜色
		EncodeCaller:   zapcore.FullCallerEncoder,      // 调用者编码器，使用完整路径
		EncodeDuration: zapcore.SecondsDurationEncoder, // 持续时间编码器，使用秒数
	}
	config.TimeKey = ""
	config.LevelKey = ""
	return zapcore.NewConsoleEncoder(config)
}

func (l *Logger) WithOption(o opt.Option) {
	l.rw.Lock()
	defer l.rw.Unlock()
	l.Opt = o
}

func (l *Logger) Debug(msgID, format string, args ...any) {
	l.rw.RLock()
	defer l.rw.RUnlock()
	l.l.Debug(l.Opt.FormatMessage(msgID, DEBUG, format, args))
}

func (l *Logger) Info(msgID, format string, args ...any) {
	l.rw.RLock()
	defer l.rw.RUnlock()
	l.l.Info(l.Opt.FormatMessage(msgID, INFO, format, args))
}

func (l *Logger) Warn(msgID, format string, args ...any) {
	l.rw.RLock()
	defer l.rw.RUnlock()
	l.l.Warn(l.Opt.FormatMessage(msgID, WARN, format, args))
}

func (l *Logger) Error(msgID, format string, args ...any) {
	l.rw.RLock()
	defer l.rw.RUnlock()
	l.l.Error(l.Opt.FormatMessage(msgID, ERROR, format, args))
}

func (l *Logger) Panic(msgID, format string, args ...any) {
	l.rw.RLock()
	defer l.rw.RUnlock()
	l.l.Panic(l.Opt.FormatMessage(msgID, PANIC, format, args))
}

func (l *Logger) Fatal(msgID, format string, args ...any) {
	l.rw.RLock()
	defer l.rw.RUnlock()
	l.l.Fatal(l.Opt.FormatMessage(msgID, FATAL, format, args))
}

func (l *Logger) Level() Level {
	return Level(l.l.Level())
}

func (l *Logger) LevelString() string {
	return l.l.Level().String()
}

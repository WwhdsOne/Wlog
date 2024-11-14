package WLog

import (
	"dev.aminer.cn/codegeex-enterprise/cgxlog/cgx-wlog-go.git/opt"
	"dev.aminer.cn/codegeex-enterprise/cgxlog/cgx-wlog-go.git/writer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

type Logger struct {
	l   *zap.Logger  // 日志实体
	Opt opt.Option   // 日志选项接口
	rw  sync.RWMutex // opt切换读写锁
}

func Build(ls *writer.WLogWriters, option opt.Option) *Logger {

	// 初始化日志输出
	writers := ls.BuildWriters()

	// 创建日志输出对象
	Wzapcore := zapcore.NewCore(
		newEncoder(),                             // 编码器
		zapcore.NewMultiWriteSyncer(writers...),  // 写入对象
		zap.NewAtomicLevelAt(zapcore.DebugLevel), // 日志级别
	)
	// 默认Rfc5424
	if option == nil {
		option = &opt.Rfc5424Opt{}
	}

	return &Logger{
		l:   zap.New(Wzapcore, zap.AddStacktrace(zapcore.ErrorLevel)),
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

func (l *Logger) Debug(msgID, msg string) {
	l.rw.RLock()
	defer l.rw.RUnlock()
	l.l.Debug(l.Opt.FormatMessage(msgID, msg, DEBUG))
}

func (l *Logger) Info(msgID, msg string) {
	l.rw.RLock()
	defer l.rw.RUnlock()
	l.l.Info(l.Opt.FormatMessage(msgID, msg, INFO))
}

func (l *Logger) Warn(msgID, msg string) {
	l.rw.RLock()
	defer l.rw.RUnlock()
	l.l.Warn(l.Opt.FormatMessage(msgID, msg, WARN))
}

func (l *Logger) Error(msgID, msg string) {
	l.rw.RLock()
	defer l.rw.RUnlock()
	l.l.Error(l.Opt.FormatMessage(msgID, msg, ERROR))
}

func (l *Logger) Panic(msgID, msg string) {
	l.rw.RLock()
	defer l.rw.RUnlock()
	l.l.Panic(l.Opt.FormatMessage(msgID, msg, PANIC))
}

func (l *Logger) Fatal(msgID, msg string) {
	l.rw.RLock()
	defer l.rw.RUnlock()
	l.l.Fatal(l.Opt.FormatMessage(msgID, msg, FATAL))
}

package test

import (
	"WLog/core"
	"WLog/writer/file"
	"WLog/writer/mq"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"testing"
)

func TestSummaryLogger(t *testing.T) {
	c := core.LogFormatConfig{
		Level:           zapcore.DebugLevel,
		Prefix:          "[TEST-ZAP-JSON]",
		IsJson:          true,
		EncoderLevel:    "LowercaseLevelEncoder",
		StacktraceLevel: zapcore.ErrorLevel,
	}

	f := &file.LocalFileLogWriter{
		FileName:    "test.log",
		FileDirPath: "/Users/wwhds/Programming_Learning/Project/zapLog/test/log",
	}

	k := &mq.KafkaLogProducer{
		Topic:     "test-topic",
		Partition: 0,
		Host:      "localhost",
		Port:      9092,
	}
	// 创建一个测试日志记录器
	l := &core.LogSummary{
		LocalFileWriter: f,
		KafkaWriter:     k,
		LogFormatConfig: &c,
	}

	summaryLogger := core.Build(l)

	// 测试 Debug 方法
	t.Run("TestDebug", func(t *testing.T) {
		summaryLogger.Debug("Debug message")
	})

	// 测试 Info 方法
	t.Run("TestInfo", func(t *testing.T) {
		summaryLogger.Info("Info message")
	})

	// 测试日志切割
	t.Run("TestSplit", func(t *testing.T) {
		for i := 0; i < 20; i++ {
			summaryLogger.Info("Info message")
		}
	})

	// 测试 Warn 方法
	t.Run("TestWarn", func(t *testing.T) {
		summaryLogger.Warn("Warn message")
	})

	// 测试 Error 方法
	t.Run("TestError", func(t *testing.T) {
		summaryLogger.Error("Error message")
	})

	// 测试 Panic 方法
	t.Run("TestPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Panic did not panic")
			}
		}()
		summaryLogger.Panic("Panic message")
	})

	// 测试 Fatal 方法
	t.Run("TestFatal", func(t *testing.T) {
		// Fatal 方法会终止程序，因此我们使用 zaptest.NewLogger 来捕获日志
		zaptest.NewLogger(t)
		summaryLogger.Fatal("Fatal message")
	})

	// 测试 Sync 方法
	t.Run("TestSync", func(t *testing.T) {
		err := summaryLogger.Sync()
		if err != nil {
			t.Errorf("Sync failed: %v", err)
		}
	})
}

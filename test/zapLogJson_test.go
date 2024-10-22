package test

import (
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"testing"
	"zapLog/core"
)

func TestJSONLogger(t *testing.T) {

	c := core.LogFormatConfig{
		Level:           zapcore.DebugLevel,
		Prefix:          "[TEST-ZAP-JSON]",
		IsJson:          true,
		EncoderLevel:    "LowercaseLevelEncoder",
		StacktraceLevel: zapcore.ErrorLevel,
	}
	// 创建一个测试日志记录器
	JsonLogger := core.Build(&core.LogSummary{
		LogFormatConfig: &c,
	})

	// 测试 Debug 方法
	t.Run("TestDebug", func(t *testing.T) {
		JsonLogger.Debug("Debug message")
	})

	// 测试 Info 方法
	t.Run("TestInfo", func(t *testing.T) {
		JsonLogger.Info("Info message")
	})

	// 测试 Warn 方法
	t.Run("TestWarn", func(t *testing.T) {
		JsonLogger.Warn("Warn message")
	})

	// 测试 Error 方法
	t.Run("TestError", func(t *testing.T) {
		JsonLogger.Error("Error message")
	})

	// 测试 Panic 方法
	t.Run("TestPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Panic did not panic")
			}
		}()
		JsonLogger.Panic("Panic message")
	})

	// 测试 Fatal 方法
	t.Run("TestFatal", func(t *testing.T) {
		// Fatal 方法会终止程序，因此我们使用 zaptest.NewLogger 来捕获日志
		zaptest.NewLogger(t)
		JsonLogger.Fatal("Fatal message")
	})

	// 测试 Sync 方法
	t.Run("TestSync", func(t *testing.T) {
		err := JsonLogger.Sync()
		if err != nil {
			t.Errorf("Sync failed: %v", err)
		}
	})
}

# WLog

WLog is a convenient and easy-to-use secondary encapsulation of ZapLog, providing different encoding formats, supporting message queues, file logging and syslog.

## Requirement

- Go: >=1.22.1
- OS: Linux / MacOS / Windows

## Features

- Provides encapsulation for various settings, allowing simple configuration to complete formatted logging.
- Offers default log rotation functionality.
- Supports message queues (currently only supports sending to a specified topic in Kafka).

# Quick Start

**First:**

```bash
go get -u github.com/WwhdsOne/Wlog@v1.0.0
go get -u go.uber.org/zap 
go get -u github.com/natefinch/lumberjack 
go get -u github.com/IBM/sarama 
```

**Then:**

```go
func main() {
	WLog.Info("LOL")
	WLog.Warn("LOL")
	WLog.Debug("LOL")
}
```

**Output:**

```bash
{"level":"info","time":"2024-10-24 15:51:26.703","msg":"[/tmp/GoLand/___go_build_WLogTest] | LOL"}
{"level":"warn","time":"2024-10-24 15:51:26.704","msg":"[/tmp/GoLand/___go_build_WLogTest] | LOL"}
{"level":"debug","time":"2024-10-24 15:51:26.704","msg":"[/tmp/GoLand/___go_build_WLogTest] | LOL"}
```

# Usage

## Direct Usage

```go
func main() {
	l := WLog.Default()
	lo := &wlcore.Loptions{
		Package: "testPackage",
		Option:  []any{"LOL", 123},
	}
	l.Debug("Debug %s %d", lo)
	l.Info("Info %s %d", lo)
	l.Warn("Warn %s %d", lo)
}
```

## Specify File Object

```go
func main() {
	ls := &wlcore.LogSummary{
		LocalFileWriter: &file.LocalFileLogWriter{FileName: "app.log", FileDirPath: "./logs"},
	}
	fmt.Println(ls)
	lo := &wlcore.Loptions{
		Package: "testPackage",
		Option:  []any{"LOL", 123},
	}
	build := WLog.Build(ls)
	build.Info("Hello World %s %d", lo)
	build.Error("Hello Error %s %d", lo)
	build.Debug("Hello Debug %s %d", lo)
}
```

## Specify Kafka and Topic

```go
func main() {
	ls := &wlcore.LogSummary{
		KafkaWriter: &mq.KafkaLogProducer{
			Topic:     "test-topic",
			Partition: 0,
			Host:      "localhost",
			Port:      9092,
		},
	}
	fmt.Println(ls)
	lo := &wlcore.Loptions{
		Package: "testPackage",
		Option:  []any{"LOL", 123},
	}
	build := WLog.Build(ls)
	build.Info("Hello World %s %d", lo)
	build.Error("Hello Error %s %d", lo)
	build.Debug("Hello Debug %s %d", lo)
}
```

## Specify syslog

```go
func main() {
	ls := &wlcore.LogSummary{
		SysLogWriter: &sys.SyslogWriter{
			Port:    515,
			Host:    "47.93.83.136",
			Network: "udp",
			Tag:     "test",
		},
	}
	lo := &wlcore.Loptions{
		Package: "testPackage",
		Option:  []any{"LOL", 123},
	}
	build := WLog.Build(ls)
	build.Info("Hello World %s %d", lo)
	build.Error("Hello Error %s %d", lo)
	build.Debug("Hello Debug %s %d", lo)
}
```

## Modify Log Format

```go
func main() {
	ls := &wlcore.LogSummary{
		LogFormatConfig: &wlcore.LogFormatConfig{
			Level:           WLog.DebugLevel,
			Prefix:          "TEST-ZAP-JSON",
			IsJson:          false,
			EncoderLevel:    WLog.CapitalColorLevelEncoder,
			StacktraceLevel: WLog.ErrorLevel,
		},
	}
	lo := &wlcore.Loptions{
		Package: "testPackage",
		Option:  []any{"LOL", 123},
	}
	build := WLog.Build(ls)
	build.Info("Hello World %s %d", lo)
	build.Error("Hello Error %s %d", lo)
	build.Debug("Hello Debug %s %d", lo)
}
```

## Replace Default WLog

```go
func main() {
	ls := &wlcore.LogSummary{
		LogFormatConfig: &wlcore.LogFormatConfig{
			Level:           WLog.DebugLevel,
			Prefix:          "TEST-ZAP-JSON",
			IsJson:          false,
			EncoderLevel:    WLog.CapitalColorLevelEncoder,
			StacktraceLevel: WLog.ErrorLevel,
		},
	}
	newDefaultLogger := WLog.Build(ls)
	WLog.ReplaceDefault(newDefaultLogger)
	WLog.Info("LOL")
	WLog.Warn("LOL")
	WLog.Debug("LOL")
}
```

# Result

```bash

> Default
2024-10-25 18:29:26.917 DEBUG   [TEST-ZAP-JSON] | Debug message
2024-10-25 18:29:26.918 INFO    [TEST-ZAP-JSON] | package = test | Info message LOL 123
2024-10-25 18:29:26.918 WARN    [TEST-ZAP-JSON] | package = test | Warn message LOL 123
2024-10-25 18:29:26.918 ERROR   [TEST-ZAP-JSON] | package = test | Error messageLOL 123

> Json
{"level":"debug","time":"2024-10-25 18:30:29.731","msg":"[TEST-ZAP-JSON] | Debug message"}
{"level":"info","time":"2024-10-25 18:30:29.732","msg":"[TEST-ZAP-JSON] | package = test | Info message LOL 123"}
{"level":"warn","time":"2024-10-25 18:30:29.732","msg":"[TEST-ZAP-JSON] | package = test | Warn message LOL 123"}

```

# Tips

If you try to use placeholders like this, you should ensure that your arguments match the number and format of the options in the `log`.

## Example

When using placeholders, you should ensure that the number and format of the arguments match the placeholders:

```go
loptions := wlcore.Loptions{
  Option: []any{1, "LOL"},
}
newDefaultLogger.Info("LOL %d %s",&loptions)
```

In this example, `%s` corresponds to the string `"LOL"`, and `%d` corresponds to the integer `1`.

> If you try to use placeholders like this, you should ensure that your arguments match the number and format of the options in the `log`.
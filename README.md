# WLog

WLog is a convenient and easy-to-use secondary encapsulation of ZapLog, providing different encoding formats, supporting message queues, and file logging.

## Requirement

- Go: >=1.22.1
- OS: Linux / MacOS / Windows
- CPU: AMD64 / ARM64

## Features

- Provides encapsulation for various settings, allowing simple configuration to complete formatted logging.
- Offers default log rotation functionality.
- Supports message queues (currently only supports sending to a specified topic in Kafka).

# Getting WLog

```bash
go get -u github.com/WwhdsOne/Wlog
go get -u go.uber.org/zap 
go get -u github.com/natefinch/lumberjack 
go get -u github.com/IBM/sarama 
```

# Effect Demonstration

```bash
{"level":"debug","time":"2024-10-22 14:10:39.550","msg":"Hello Debug"}
{"level":"info","time":"2024-10-22 14:10:39.549","msg":"Hello World"}
{"level":"warn","time":"2024-10-22 14:15:42.913","msg":"Warn message"}
{"level":"error","time":"2024-10-22 14:15:42.913","msg":"Error message","stacktrace":"github.com/WwhdsOne/Wlog/WLog.(*Logger).Error/WLog/WLog/WLog.go:85\ncommand-line-arguments.TestJSONLogger.func4/WLog/test/WLogJson_test.go:41\ntesting.tRunner\n\t/Users/wwhds/go/go1.22.1/src/testing/testing.go:1689"}
{"level":"panic","time":"2024-10-22 14:15:42.914","msg":"Panic message","stacktrace":"github.com/WwhdsOne/Wlog/WLog.(*Logger).Panic/WLog/WLog/WLog.go:89\ncommand-line-arguments.TestJSONLogger.func5/WLog/test/WLogJson_test.go:51\ntesting.tRunner\n\t/Users/wwhds/go/go1.22.1/src/testing/testing.go:1689"}
```

## Usage

### Direct Usage

```go
func main() {
	logger := WLog.Default()
	logger.Info("Hello World")
	logger.Debug("Hello World")
	logger.Error("Hello World")
}
```

### Specify File Object

```go
ls := &WLog.LogSummary{
  LocalFileWriter: &file.LocalFileLogWriter{FileName: "app.log", FileDirPath: "./logs"},
}
fmt.Println(ls)

build := WLog.Build(ls)
build.Info("Hello World")
build.Error("Hello Error")
build.Debug("Hello Debug")
```

### Specify Kafka and Topic

```go
ls := &WLog.LogSummary{
  KafkaWriter:     &mq.KafkaLogProducer{Topic: "test-topic", 
                                        Partition: 0, 
                                        Host: "localhost", 
                                        Port: 9092
                                       },
}
fmt.Println(ls)

build := WLog.Build(ls)
build.Info("Hello World")
build.Error("Hello Error")
build.Debug("Hello Debug")
```

### Modify Log Format

```go
ls := &WLog.LogSummary{
  LocalFileWriter: &file.LocalFileLogWriter{FileName: "app.log", FileDirPath: "./logs"},
  LogFormatConfig: &WLog.LogFormatConfig{
    Level:           zapcore.DebugLevel,
    Prefix:          "[TEST-ZAP-JSON]",
    IsJson:          false,
    EncoderLevel:    "CapitalColorLevelEncoder",
    StacktraceLevel: zapcore.ErrorLevel,
  },
}
fmt.Println(ls)

build := WLog.Build(ls)
build.Info("Hello World")
build.Error("Hello Error")
build.Debug("Hello Debug")
```

### Result

```bash
[TEST-ZAP-JSON] 2024-10-22 14:12:35.517 INFO    Hello World
[TEST-ZAP-JSON] 2024-10-22 14:12:35.519 ERROR   Hello Error
```
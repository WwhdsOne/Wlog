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

## Usage

### Direct Usage

```go
func main() {
  l := WLog.Default()
  lo := &WLog.Loptions{
    Package: "testPackage",
    Option:  []any{"LOL", 123},
  }
  l.Debug("Debug %s %d", lo)
  l.Info("Info %s %d", lo)
  l.Warn("Warn %s %d", lo)
}
```

### Specify File Object

```go
ls := &WLog.LogSummary{
  LocalFileWriter: &file.LocalFileLogWriter{FileName: "app.log", FileDirPath: "./logs"},
}
fmt.Println(ls)
lo := &WLog.Loptions{
    Package: "testPackage",
    Option:  []any{"LOL", 123},
  }
build := WLog.Build(ls)
build.Info("Hello World",lo)
build.Error("Hello Error",lo)
build.Debug("Hello Debug",lo)
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
lo := &WLog.Loptions{
  Package: "testPackage",
  Option:  []any{"LOL", 123},
}
build := WLog.Build(ls)
build.Info("Hello World",lo)
build.Error("Hello Error",lo)
build.Debug("Hello Debug",lo)
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
lo := &WLog.Loptions{
    Package: "testPackage",
    Option:  []any{"LOL", 123},
  }
build := WLog.Build(ls)
build.Info("Hello World",lo)
build.Error("Hello Error",lo)
build.Debug("Hello Debug",lo)
```

### Result

```bash

> JSON
{"level":"info","time":"2024-10-23 14:14:48.293","msg":"[TEST-ZAP-JSON] package = test Info message LOL 123"}
Message sent to partition 0 at offset 20116
{"level":"info","time":"2024-10-23 14:14:48.295","msg":"[TEST-ZAP-JSON] package = test Info message LOL 123"}

> Default
2024-10-23 14:30:29.550 DEBUG   [TEST-ZAP-JSON] Debug message
2024-10-23 14:30:29.550 INFO    [TEST-ZAP-JSON] package = test Info message LOL 123
2024-10-23 14:30:29.550 WARN    [TEST-ZAP-JSON] package = test Warn message LOL 123
```
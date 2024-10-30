# WLog

**WLog** 是一个便捷且易于使用的 ZapLog 二次封装，WLog 可以使用 RFC5424格式，支持上报消息队列、文件日志和 syslog。以满足更广泛的日志记录需求。

## 环境需求

- Go: >=1.22.1
- OS: Linux / MacOS / Windows

## 特性

- **多种设置封装**：提供对各种设置的封装，允许通过简单的配置完成格式化日志记录。
- **默认日志轮转功能**：提供默认的日志轮转功能。
- **支持消息队列**：支持消息队列（目前仅支持发送到 Kafka 中的指定主题）。

# 快速开始

**首先**

```bash
go get -u github.com/WwhdsOne/Wlog
```

**新建一个日志并调用方法**

```go
func main() {
  // 创建一个测试日志记录器
  DefaultLogger := WLog.Build(&writer.WLogWriters{}, nil)
  // 打印日志
  DefaultLogger.Debug("233", "Debug message")
  DefaultLogger.Warn("233", "Warn message")
  DefaultLogger.Info("233", "Info message")
}
```

**控制台打印**

```bash
<15>1 2024-10-30T18:58:35.901214+08:00 wangwenhaideMacBook-Air.local eex_enterprise_wlog_test__TestDefaultLogger.test 67231 233 [] Debug message
<14>1 2024-10-30T18:58:35.902344+08:00 wangwenhaideMacBook-Air.local eex_enterprise_wlog_test__TestDefaultLogger.test 67231 233 [] Info message
<12>1 2024-10-30T18:58:35.902377+08:00 wangwenhaideMacBook-Air.local eex_enterprise_wlog_test__TestDefaultLogger.test 67231 233 [] Warn message
```

# 使用

## 指定文件对象

```go
f := &writer.LocalFileLogWriter{
  FileName:    "test.log",
  FileDirPath: "./log",
}

lumberjackLogger := WLog.Build(&writer.WLogWriters{
  LocalFileWriter: f,
}, nil)
```

文件对象自带日志切割，属性如下

```go
lumberJackLogger := &lumberjack.Logger{
  Filename:   filename, // 文件位置
  MaxSize:    10,       // 进行切割之前,日志文件的最大大小(MB为单位)
  MaxAge:     7,        // 保留旧文件的最大天数
  MaxBackups: 10,       // 保留旧文件的最大个数
  Compress:   false,    // 是否压缩/归档旧文件
  LocalTime:  true,     // 是否使用本地时间
}
```

## 指定Kafka和Topic

```go
k := &writer.KafkaLogProducer{
  Topic:     "test-topic",
  Partition: 0,
  Host:      "localhost",
  Port:      9092,
}
// 创建一个测试日志记录器
l := &writer.WLogWriters{
  KafkaWriter: k,
}
KafkaLogger := WLog.Build(l, nil)
```

## 指定Syslog

```go
s := &writer.SyslogWriter{
  Host:     "47.93.83.136",
  Port:     515,
  Network:  "udp",
  Priority: syslog.LOG_INFO,
  Tag:      "WWh",
}

// 创建一个测试日志记录器
r := opt.Rfc5424Opt{
  //AppName: "/var/folders/4c/j8r4jbh539s_gkfxssjw9dtm0000gp/T/go-build2659986742/b001/exe/main",
}
r.AddDatum("233", "777", "666")
Rfc5424Logger := WLog.Build(&writer.WLogWriters{
  SysLogWriter: s,
}, &r)
```

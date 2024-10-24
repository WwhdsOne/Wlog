package file

import (
	"github.com/natefinch/lumberjack"
	"io"
	"path"
)

type LocalFileLogWriter struct {
	FileDirPath string    //文件目录
	FileName    string    //文件名
	Writer      io.Writer // 日志写入器
}

// InitWriter 创建一个日志切割对象
func (l *LocalFileLogWriter) InitWriter() {
	// 获取文件名称
	filename := path.Join(l.FileDirPath, l.FileName)
	// 创建一个日志切割对象
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename, // 文件位置
		MaxSize:    10,       // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     7,        // 保留旧文件的最大天数
		MaxBackups: 10,       // 保留旧文件的最大个数
		Compress:   false,    // 是否压缩/归档旧文件
		LocalTime:  true,     // 是否使用本地时间
	}
	l.Writer = lumberJackLogger
}

func (l *LocalFileLogWriter) Write(p []byte) (n int, err error) {
	return l.Writer.Write(p)
}

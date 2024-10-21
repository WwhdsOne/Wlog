package file

import (
	"github.com/natefinch/lumberjack"
	"io"
	"path"
)

type LocalFileLogWriter struct {
	FileDirPath string `yaml:"fileDirPath"` //文件目录
	FileName    string `yaml:"fileName"`    //文件名
}

// CreateWriter 创建一个日志切割对象
func (l *LocalFileLogWriter) CreateWriter() io.Writer {
	// 获取文件名称
	filename := path.Join(l.FileDirPath, l.FileName)
	// 创建一个日志切割对象
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename, // 文件位置
		MaxSize:    1,        // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     7,        // 保留旧文件的最大天数
		MaxBackups: 10,       // 保留旧文件的最大个数
		Compress:   false,    // 是否压缩/归档旧文件
		LocalTime:  true,     // 是否使用本地时间
	}
	return lumberJackLogger
}

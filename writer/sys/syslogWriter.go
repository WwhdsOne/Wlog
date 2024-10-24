package sys

import (
	"github.com/WwhdsOne/Wlog"
	"log/syslog"
	"strconv"
)

type SyslogWriter struct {
	Network   string          //网络通信方式
	Host      string          // ip地址
	Port      int             // 端口
	Priority  syslog.Priority // 优先级
	Tag       string          // 标签
	SysWriter *syslog.Writer  // syslog writer
}

// InitWriter 初始化SysWriter
func (s *SyslogWriter) InitWriter() {
	writer, err := syslog.Dial(s.Network, s.Host+":"+strconv.Itoa(s.Port), s.Priority, s.Tag)
	if err != nil {
		WLog.Warn(err.Error())
	}
	s.SysWriter = writer
}

// 提供io.Writer
func (s *SyslogWriter) Write(p []byte) (n int, err error) {
	return s.SysWriter.Write(p)
}

package writer

import (
	"log"
	"log/syslog"
	"net"
)

type SyslogWriter struct {
	Network  string          //网络通信方式
	Host     string          // ip地址
	Port     int             // 端口
	Priority syslog.Priority // 优先级
	Tag      string          // 标签
	conn     *net.UDPConn
}

// InitWriter 初始化SysWriter
func (s *SyslogWriter) InitWriter() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP(s.Host),
		Port: s.Port,
	})
	if err != nil {
		log.Printf("Failed to connect to syslog server: %s\n", err)

	}
	s.conn = conn
}

func (s *SyslogWriter) Write(p []byte) (n int, err error) {
	return s.conn.Write(p)
}

func (s *SyslogWriter) Close() error {
	return s.conn.Close()
}

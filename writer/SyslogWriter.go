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
	conn     net.Conn
}

// InitWriter 初始化SysWriter
func (s *SyslogWriter) InitWriter() {
	var conn net.Conn
	var err error
	// 使用裸的udp和tcp协议发送，dial会自己包上一层使得日志无法被正常解析
	if s.Network == "udp" {
		conn, err = net.DialUDP(s.Network, nil, &net.UDPAddr{
			IP:   net.ParseIP(s.Host),
			Port: s.Port,
		})
	} else if s.Network == "tcp" {
		conn, err = net.DialTCP(s.Network, nil, &net.TCPAddr{
			IP:   net.ParseIP(s.Host),
			Port: s.Port,
		})
	} else {
		log.Printf("Network type %s is not supported\n", s.Network)
		return
	}
	if err != nil {
		log.Printf("Failed to connect to syslog server: %s\n", err)
		return
	}
	s.conn = conn
}

func (s *SyslogWriter) Write(p []byte) (n int, err error) {
	return s.conn.Write(p)
}

func (s *SyslogWriter) Close() error {
	return s.conn.Close()
}

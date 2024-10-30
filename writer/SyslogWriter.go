package writer

import (
	"fmt"
	"log"
	"log/syslog"
	"net"
)

func init() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(10, 50, 34, 72),
		Port: 51898,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	fmt.Println(string(b))
	x, err := socket.Write(b)
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	fmt.Println("发送成功，发送字节数：", x)
}

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

package wlcore

type Loptions struct {
	Package        string          // 包名
	Option         []any           // 其他需要打印的选项
	Rfc5424Message *Rfc5424Message // rfc5424信息
}

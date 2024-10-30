package writer

type logWriter interface {
	InitWriter()                       // InitWriter 初始化Writer
	Write(p []byte) (n int, err error) // Write 实现io.Writer接口
}

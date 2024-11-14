package opt

type Option interface {
	//FormatMessage(msgID string, lv int, msg string) string
	FormatMessage(msgID string, lv int, format string, args ...any) string
}

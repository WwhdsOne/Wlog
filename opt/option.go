package opt

type Option interface {
	FormatMessage(msgID, msg string, lv int) string
}

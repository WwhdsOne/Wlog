package opt

import "context"

type Option interface {
	//FormatMessage(msgID string, lv int, msg string) string
	FormatMessage(msgID string, lv int, format string, args []any) string
	WithContext(ctx context.Context)
	WithContextKeys(keys []any)
}

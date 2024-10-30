package wlcore

import (
	"fmt"
	"strings"
	"time"
)

type Loptions struct {
	Package        string          // 包名
	Option         []any           // 其他需要打印的选项
	Rfc5424Message *Rfc5424Message // rfc5424信息
}

// FormatRfc5424Message formats the given parameters into an RFC 5424 compliant log message
func (l *Loptions) FormatRfc5424Message(hostname, appname, msg string) string {
	m := l.Rfc5424Message
	if m == nil {
		return fmt.Sprintf("<%d>1 %s %s %s %s %s %s %s",
			6+8, // 假设为Info
			time.Now().UTC().Format(time.RFC3339),
			hostname,
			appname,
			"-",
			"-",
			"[]",
			msg,
		)
	}
	// Calculate the PRI value
	pri := int(m.Priority) + 8 // Assuming facility is user (1)

	// Get the current timestamp in RFC 5424 format
	timestamp := time.Now().UTC().Format(time.RFC3339)

	// Replace empty parameters with "-"
	if hostname == "" {
		hostname = "-"
	}
	if appname == "" {
		appname = "-"
	}
	if m.ProcessID == "" {
		m.ProcessID = "-"
	}
	if m.MessageID == "" {
		m.MessageID = "-"
	}

	// Format the structured data
	structuredData := formatStructuredData(m.StructuredData)
	if structuredData == "" {
		structuredData = "-"
	}

	// Construct the final log m
	logMessage := fmt.Sprintf("<%d>1 %s %s %s %s %s %s %s",
		pri,
		timestamp,
		hostname,
		appname,
		m.ProcessID,
		m.MessageID,
		structuredData,
		msg,
	)
	return logMessage
}

func (l *Loptions) SetRfcLevel(level int) {
	rm := l.Rfc5424Message
	if rm != nil {
		rm.Priority = level
	}
}

// formatStructuredData formats the structured data into the RFC 5424 format
func formatStructuredData(data []StructuredData) string {
	if len(data) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, sd := range data {
		sb.WriteString(fmt.Sprintf("[%s", sd.ID))
		for _, param := range sd.Parameters {
			sb.WriteString(fmt.Sprintf(" %s=\"%s\"", param.Name, param.Value))
		}
		sb.WriteString("]")
	}

	return sb.String()
}

package wlcore

import (
	"fmt"
	"strings"
	"time"
)

type Priority int

const (
	Emergency Priority = iota
	Alert
	Crit
	Error
	Warning
	Notice
	Info
	Debug
)

type Rfc5424Message struct {
	Priority       Priority
	ProcessID      string
	MessageID      string
	StructuredData []StructuredData
}

// SDParam represents parameters for structured data
type SDParam struct {
	Name  string
	Value string
}

// StructuredData represents structured data within a log message
type StructuredData struct {
	ID         string
	Parameters []SDParam
}

// AddDatum adds structured data to a log message
func (m *Rfc5424Message) AddDatum(ID string, Name string, Value string) {
	if m.StructuredData == nil {
		m.StructuredData = []StructuredData{}
	}
	for i, sd := range m.StructuredData {
		if sd.ID == ID {
			sd.Parameters = append(sd.Parameters, SDParam{Name: Name, Value: Value})
			m.StructuredData[i] = sd
			return
		}
	}

	m.StructuredData = append(m.StructuredData, StructuredData{
		ID: ID,
		Parameters: []SDParam{
			{
				Name:  Name,
				Value: Value,
			},
		},
	})
}

// FormatRfc5424Message formats the given parameters into an RFC 5424 compliant log message
func (m *Rfc5424Message) FormatRfc5424Message(hostname, appname, msg string) string {
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

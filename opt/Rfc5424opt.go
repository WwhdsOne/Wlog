package opt

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	Emergency = iota
	Alert
	Crit
	Error
	Warning
	Notice
	Info
	Debug
)

type Rfc5424Opt struct {
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

func convertLogLevel(customLevel int) int {
	switch customLevel {
	case 0:
		return Debug
	case 1:
		return Info
	case 2:
		return Warning
	case 3:
		return Error
	case 4:
		return Crit
	case 5:
		return Emergency
	default:
		return -1 // 未知级别
	}
}

// AddDatum adds structured data to a log message
func (r *Rfc5424Opt) AddDatum(ID string, Name string, Value string) {
	if r.StructuredData == nil {
		r.StructuredData = []StructuredData{}
	}
	for i, sd := range r.StructuredData {
		if sd.ID == ID {
			sd.Parameters = append(sd.Parameters, SDParam{Name: Name, Value: Value})
			r.StructuredData[i] = sd
			return
		}
	}

	r.StructuredData = append(r.StructuredData, StructuredData{
		ID: ID,
		Parameters: []SDParam{
			{
				Name:  Name,
				Value: Value,
			},
		},
	})
}

// formatStructuredData formats the structured data into the RFC 5424 format
func (r *Rfc5424Opt) formatStructuredData() string {
	data := r.StructuredData
	if len(data) == 0 {
		return "[]"
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

// FormatMessage formats the given parameters into an RFC 5424 compliant log message
func (r *Rfc5424Opt) FormatMessage(msgID, msg string, lv int) string {
	hostname, _ := os.Hostname()

	// Calculate the PRI value
	pri := convertLogLevel(lv) + 8 // Assuming the facility is always user (1)

	// Get the current timestamp in RFC 5424 format
	timestamp := time.Now().UTC().Format(time.RFC3339)

	if msgID == "" {
		msgID = "-"
	}

	// Format the structured data
	structuredData := r.formatStructuredData()

	// Construct the final log m
	logMessage := fmt.Sprintf("<%d>1 %s %s %s %d %s %s %s",
		pri,
		timestamp,
		hostname,
		os.Args[0],
		os.Getgid(),
		msgID,
		structuredData,
		msg,
	)
	return logMessage
}

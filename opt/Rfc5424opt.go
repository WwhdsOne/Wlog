package opt

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

const RFC3339Micro = "2006-01-02T15:04:05.999999Z07:00"

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
	Hostname, AppName string
	StructuredData    []StructuredData
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

// SetDatum sets structured data to a log message
func (r *Rfc5424Opt) SetDatum(ID string, Name string, Value string) {
	for i, sd := range r.StructuredData {
		if sd.ID == ID {
			for j, param := range sd.Parameters {
				if param.Name == Name {
					sd.Parameters[j].Value = Value
					r.StructuredData[i] = sd
					return
				}
			}
			sd.Parameters = append(sd.Parameters, SDParam{Name: Name, Value: Value})
			r.StructuredData[i] = sd
			return
		}
	}
	r.AddDatum(ID, Name, Value)
}

// formatStructuredData formats the structured data into the RFC 5424 format
func (r *Rfc5424Opt) formatStructuredData() string {
	data := r.StructuredData
	if len(data) == 0 {
		return "-"
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

func (r *Rfc5424Opt) GetAppName() string {
	// 如果appName为空字符串的话，使用os.Args[0]作为默认值
	if r.AppName == "" {
		appName := os.Args[0]
		// 如果appName中包含路径的话，那么就截取路径后面的部分
		if start := strings.LastIndex(appName, "/"); start != -1 {
			appName = appName[start+1:]
		}
		r.AppName = appName
	}

	// 如果长度大于48那么进行截取
	if len(r.AppName) > 48 {
		r.AppName = r.AppName[len(r.AppName)-48:]
	}

	return r.AppName
}

func (r *Rfc5424Opt) GetHostname() string {
	if r.Hostname == "" {
		r.Hostname, _ = os.Hostname()
	}
	return r.Hostname
}

func (r *Rfc5424Opt) FormatMessage(msgID string, lv int, format string, args ...any) string {

	// Calculate the PRI value
	pri := convertLogLevel(lv) + 8 // Assuming the facility is always user (1)

	// Get the current timestamp in RFC 5424 format
	timestamp := time.Now().Format(RFC3339Micro)

	if msgID == "" {
		msgID = "-"
	}

	msg := "-"
	if format != "" {
		msg = format
		fmt.Println(args)
		fmt.Println(args[0], reflect.TypeOf(args[0]))
		if len(args) > 0 && args[0] != nil {
			msg = fmt.Sprintf(format, args...)
		}
	}

	// Format the structured data
	structuredData := r.formatStructuredData()

	// Construct the final log m
	logMessage := fmt.Sprintf("<%d>1 %s %s %s %d %s %s %s\n",
		pri,
		timestamp,
		r.GetHostname(),
		r.GetAppName(),
		os.Getpid(),
		msgID,
		structuredData,
		msg,
	)
	return logMessage
}

package wlcore

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

type Rfc5424Message struct {
	Priority       int
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

package log

type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
)

// String returns the string representation of the log level.
func (level Level) String() string {
	switch level {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}

}

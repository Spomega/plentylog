package log

// Level represents the log level.
type Level int

const (
	// Debug level.
	Debug Level = iota
	// Info level.
	Info
	// Warn level.
	Warn
	// Error level.
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

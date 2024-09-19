package domain

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

func (level LogLevel) String() string {
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

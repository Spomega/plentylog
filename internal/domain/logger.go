package log

// Driver represents a log driver.
type Driver interface {
	WriteLog(record *Record) error
}

// Logger represents a logger.
type Logger struct {
	drivers []Driver
}

// NewLogger creates a new logger.
func NewLogger(drivers ...Driver) *Logger {
	return &Logger{drivers: drivers}
}

// Log sends logs to all drivers.
func (logger *Logger) Log(level Level, message string, tags map[string]string, transactionID string) {
	record := NewRecord(level, message, tags, transactionID)
	for _, driver := range logger.drivers {
		_ = driver.WriteLog(record)
	}
}

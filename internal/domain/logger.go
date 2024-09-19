package log

type Driver interface {
	WriteLog(record *Record) error
}

type Logger struct {
	drivers []Driver
}

func NewLogger(drivers ...Driver) *Logger {
	return &Logger{drivers: drivers}
}

func (logger *Logger) Log(level Level, message string, tags map[string]string, transactionID string) {
	record := NewRecord(level, message, tags, transactionID)
	for _, driver := range logger.drivers {
		_ = driver.WriteLog(record)
	}
}

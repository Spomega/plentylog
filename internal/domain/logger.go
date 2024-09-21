package log

import (
	"fmt"
	"sync"
)

// Driver represents a log driver.
type Driver interface {
	WriteLog(record *Record) error
}

// Logger represents a logger.
type Logger struct {
	drivers []Driver
	mu      sync.Mutex
}

// NewLogger creates a new logger.
func NewLogger() *Logger {
	return &Logger{
		drivers: make([]Driver, 0),
	}
}

// AddDriver adds a driver to the logger.
func (l *Logger) AddDriver(driver Driver) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.drivers = append(l.drivers, driver)
}

// Log sends logs to all drivers.
func (l *Logger) Log(level Level, message string, tags map[string]string, transactionID string) {
	record := NewRecord(level, message, tags, transactionID)
	for _, driver := range l.drivers {
		err := driver.WriteLog(record)
		if err != nil {
			fmt.Printf("Error writing log: %v ", err)
			continue
		}
	}
}

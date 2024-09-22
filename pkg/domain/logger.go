package log

import (
	"fmt"
	"sync"
)

// Driver represents a log driver.
type Driver interface {
	WriteLog(record *Record) error
}

// Closable is the interface for drivers that need to be closed.
type Closable interface {
	Close() error
}

// Logger represents a logger.
type Logger struct {
	drivers   []Driver
	closables []Closable
	mu        sync.Mutex
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

	if closable, ok := driver.(Closable); ok {
		l.closables = append(l.closables, closable)
	}
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

// CloseAll closes all file-based drivers.
func (l *Logger) CloseAll() error {
	for _, closable := range l.closables {
		if err := closable.Close(); err != nil {
			return err
		}
	}
	return nil
}

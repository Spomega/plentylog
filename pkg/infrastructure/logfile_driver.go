package infrastructure

import (
	"fmt"
	log "github.com/Spomega/plentylog/pkg/domain"
	"os"
)

// LogFileDriver represents a log file driver.
type LogFileDriver struct {
	file *os.File
}

// NewLogFileDriver a log record to a file.
func NewLogFileDriver(filePath string) (*LogFileDriver, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, err
	}

	return &LogFileDriver{file: file}, nil

}

func (d *LogFileDriver) WriteLog(record *log.Record) error {
	logLine := fmt.Sprintf("[%s] %s: %s", record.Timestamp.Format("2006-01-02T15:04:05Z07:00"), record.Level, record.Message)

	if record.TransactionID != "" {
		logLine += fmt.Sprintf(" [TransactionID: %s]", record.TransactionID)
	}

	if len(record.MetaData) > 0 {
		logLine += " [Attributes:"
		for key, value := range record.MetaData {
			logLine += fmt.Sprintf(" %s=%s", key, value)
		}
		logLine += "]"
	}

	logLine += "\n"

	_, err := d.file.WriteString(logLine)
	return err
}

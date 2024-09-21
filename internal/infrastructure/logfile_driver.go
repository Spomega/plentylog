package infrastructure

import "os"

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

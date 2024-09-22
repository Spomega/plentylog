package infrastructure

import (
	"encoding/json"
	log "github.com/Spomega/plentylog/pkg/domain"
	"os"
)

// JSONFileDriver represents a JSON file driver.
type JSONFileDriver struct {
	file *os.File
}

// NewJSONFileDriver creates a new JSON file driver.
func NewJSONFileDriver(filePath string) (*JSONFileDriver, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, err
	}

	return &JSONFileDriver{file: file}, nil
}

// WriteLog writes a log record to a JSON file.
func (j *JSONFileDriver) WriteLog(record *log.Record) error {
	logRecordJSON, err := json.Marshal(record)
	if err != nil {
		return err
	}

	_, err = j.file.Write(logRecordJSON)
	_, _ = j.file.Write([]byte("\n"))

	return err
}

// Close closes the file when finished.
func (j *JSONFileDriver) Close() error {
	return j.file.Close()
}

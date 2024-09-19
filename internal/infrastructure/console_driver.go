package infrastructure

import (
	"fmt"
	log "plentylog/internal/domain"
)

type ConsoleDriver struct{}

// WriteLog writes a log record to the console.
func (c *ConsoleDriver) WriteLog(record *log.Record) error {
	fmt.Printf("[%s] %s: %s", record.Timestamp.Format("2006-01-02T15:04:05Z07:00"), record.Level.String(), record.Message)

	if record.TransactionID != "" {
		fmt.Printf(" [TransactionID: %s]", record.TransactionID)

	}

	if len(record.MetaData) > 0 {
		fmt.Printf("[Attributes]:")
		for key, value := range record.MetaData {
			fmt.Printf("%s:%s n", key, value)
		}
	}

	fmt.Println()

	return nil
}

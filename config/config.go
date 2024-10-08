package config

import (
	"encoding/json"
	"fmt"
	log "github.com/Spomega/plentylog/pkg/domain"
	"github.com/Spomega/plentylog/pkg/infrastructure"
	"os"
)

// Config represents the configuration for the logger.
type Config struct {
	Drivers []DriverConfig
}

// DriverConfig represents the configuration for a driver.
type DriverConfig struct {
	Type     string `json:"type"`
	FileName string `json:"filename"`
}

// GetLoggerWithConfig creates a new logger based on the config file.
func GetLoggerWithConfig(filePath string) (*log.Logger, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open config file: %w", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)

	if err != nil {
		return nil, fmt.Errorf("could not decode config file: %w", err)
	}

	logger, err := config.getLogger()

	if err != nil {
		return nil, fmt.Errorf("could not create logger: %w", err)
	}

	return logger, nil
}

// getLogger creates a new logger based on the config.
func (c *Config) getLogger() (*log.Logger, error) {
	logger := log.NewLogger()

	if len(c.Drivers) == 0 {
		return nil, fmt.Errorf("no drivers provided")
	}

	for _, driverConfig := range c.Drivers {

		switch driverConfig.Type {
		case "cli":
			logger.AddDriver(&infrastructure.ConsoleDriver{})
		case "json":
			if driverConfig.FileName == "" {
				return nil, fmt.Errorf("file driver requires a filename")
			}
			jsonDriver, err := infrastructure.NewJSONFileDriver(driverConfig.FileName)
			if err != nil {
				return nil, err
			}
			logger.AddDriver(jsonDriver)
		case "logfile":
			fileDriver, err := infrastructure.NewLogFileDriver(driverConfig.FileName)
			if err != nil {
				return nil, err
			}
			logger.AddDriver(fileDriver)
		default:
			return nil, fmt.Errorf("unknown driver type: %s", driverConfig.Type)
		}
	}

	return logger, nil
}

// GetDefaultLogger creates a new logger with default settings.
func GetDefaultLogger() (*log.Logger, error) {
	config := Config{
		Drivers: []DriverConfig{
			{
				Type: "cli",
			},
			{
				Type:     "json",
				FileName: "logs.json",
			},
			{
				Type:     "logfile",
				FileName: "logs.log",
			},
		},
	}

	logger, err := config.getLogger()

	if err != nil {
		return nil, fmt.Errorf("could not create logger: %w", err)
	}
	return logger, nil
}

package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func CreateLogger(levelString string) (*zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(levelString)
	if err != nil {
		return nil, err
	}
	logger := zerolog.New(os.Stdout)
	logger.Level(level)
	return &logger, nil
}

package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func InitLogger() {
	once.Do(func() {
		config := zap.NewProductionConfig()
		config.OutputPaths = []string{
			"stdout",
			"/tmp/secretsmith.log",
		}
		var err error
		logger, err = config.Build()
		if err != nil {
			panic(err)
		}
	})
}

func GetLogger() *zap.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}

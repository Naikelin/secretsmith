package logger

import (
	"log"
	"sync"

	"go.uber.org/zap"
)

var (
	once sync.Once
)

func InitLogger() *zap.Logger {
	var logger *zap.Logger
	var err error
	once.Do(func() {
		config := zap.NewProductionConfig()
		config.OutputPaths = []string{"stdout", "/tmp/secretsmith.log"}

		logger, err = config.Build()
		if err != nil {
			log.Fatal("Error building logger: ", err)
		}
	})
	return logger
}

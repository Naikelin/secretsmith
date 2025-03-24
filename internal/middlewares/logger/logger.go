package logger

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type contextKey string

const loggerKey contextKey = "logger"

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), loggerKey, logger)
		c.Request = c.Request.WithContext(ctx)
		requestId := c.GetString("RequestID")
		logger.Info("Request received", zap.String("request id", requestId), zap.String("method", c.Request.Method), zap.String("path", c.Request.URL.Path))
		c.Next()
		logger.Info("Request completed", zap.String("request id", requestId), zap.Int("status", c.Writer.Status()))
	}
}

func GetLogger(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(loggerKey).(*zap.Logger)
	if !ok {
		return zap.NewNop()
	}
	return logger
}

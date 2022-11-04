package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"log"
)

func ZapLogger(c *gin.Context) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}
	c.Next()

	body, _ := io.ReadAll(c.Request.Body)

	logger.Info("endpoint access",
		zap.String("Method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.Int("status", c.Writer.Status()),
		zap.String("request", string(body)),
	)
}

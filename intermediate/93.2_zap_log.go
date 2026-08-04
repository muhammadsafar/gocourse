package main

import (
	"log"

	"go.uber.org/zap"
)

func main932() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error in initializing Zap error")
	}

	defer logger.Sync()

	logger.Info("This is an info message from zap logger")
	logger.Info("User logged in", zap.String("username", "musa"), zap.String("method", "GET"))
	logger.Warn("This is a warning message from zap logger")
	logger.Error("This is an error message from zap logger")

}

package main

import "github.com/sirupsen/logrus"

func main931() {

	log := logrus.New()

	// Set log level
	log.SetLevel(logrus.InfoLevel)

	//set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	// Log messages
	log.Info("This is an info message")
	log.Warn("This is a warning message")
	log.Error("This is an error message")

	// Log with fields
	log.WithFields(logrus.Fields{
		"username": "musa",
		"method":   "GET",
	}).Info("User logged in")
}

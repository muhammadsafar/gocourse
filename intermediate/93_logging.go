package main

import (
	"log"
	"os"
)

func main93() {

	log.Println("This is a test log message")

	log.SetPrefix("CUSTOM LOG: ")
	log.Println("This is another log message with custom prefix")

	//log flags

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with date, time and short file info  ")

	infoLogger.Println("This is an info message")
	warnLogger.Println("This is a warning message")
	errorLogger.Println("This is an error message")

	//open file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) //0666 means read and write permission
	if err != nil {
		log.Fatalf("Failed to open log file:%v", err)
	}

	defer file.Close()

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(file, "WARN: ", log.Ldate|log.Ltime)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger.Println("This is a debug message")
	warnLogger.Println("This is a warning message written to file")
	infoLogger.Println("This is an info message written to file")
	errorLogger.Println("This is an error message written to file")

}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
)

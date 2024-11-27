package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type logEntry struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

type customLogger struct {
	fileLogger    *log.Logger
	consoleLogger *log.Logger
}

var Logger *customLogger

func New(logFile string) error {
	if Logger != nil {
		return nil
	}
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("could not open log file: %v", err)
	}
	Logger = &customLogger{
		fileLogger:    log.New(file, "", 0),
		consoleLogger: log.New(os.Stdout, "", 0),
	}
	return nil
}

func (l *customLogger) log(level string, args ...interface{}) {
	entry := &logEntry{
		Time:    time.Now().Format(time.DateTime),
		Level:   level,
		Message: strings.Join(formatArgs(args), " "),
	}

	jsonMessage, err := json.Marshal(entry)
	if err != nil {
		log.Println("ERROR: Failed to serialize log message to JSON:", err)
		return
	}

	l.fileLogger.Println(string(jsonMessage))
	l.consoleLogger.Printf("%s [%s] %s", entry.Time, entry.Level, entry.Message)
}

func (l *customLogger) Debug(args ...interface{}) {
	l.log("DEBUG", args...)
}

func (l *customLogger) Info(args ...interface{}) {
	l.log("INFO", args...)
}

func (l *customLogger) Warn(args ...interface{}) {
	l.log("WARN", args...)
}

func (l *customLogger) Error(args ...interface{}) {
	l.log("ERROR", args...)
}

func (l *customLogger) Fatal(args ...interface{}) {
	l.log("FATAL", args...)
	os.Exit(1)
}

func (l *customLogger) Panic(args ...interface{}) {
	l.log("PANIC", args...)
	panic(fmt.Sprint(args...))
}

func formatArgs(args []interface{}) []string {
	var result []string
	for _, arg := range args {
		result = append(result, fmt.Sprint(arg))
	}
	return result
}

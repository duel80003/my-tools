package tools

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	DebugLevel = "debug"
	ErrorLevel = "error"
)

var (
	Logger = logrus.New()
)

type MyTextFormatter struct{}

func (t *MyTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timeStr := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d", entry.Time.Year(), entry.Time.Month(), entry.Time.Day(), entry.Time.Hour(), entry.Time.Minute(), entry.Time.Second())
	filePath := strings.Split(entry.Caller.File, "/")
	directName := filePath[len(filePath)-2]
	name := filePath[len(filePath)-1]
	result := fmt.Sprintf("[%s] [%s/%s:%d] [%s] %s", timeStr, directName, name, entry.Caller.Line, entry.Level, entry.Message)
	return append([]byte(result), '\n'), nil
}

func LogInit() {
	Logger.SetFormatter(&MyTextFormatter{})
	Logger.Out = os.Stdout
	Logger.SetLevel(logrus.InfoLevel)
	env := os.Getenv("LOG_LEVEL")
	if strings.ToLower(env) == DebugLevel {
		Logger.SetLevel(logrus.DebugLevel)
	}
	Logger.SetReportCaller(true)
}

func LogInitWithLevel(level string) {
	Logger.SetFormatter(&MyTextFormatter{})
	Logger.Out = os.Stdout
	Logger.SetReportCaller(true)
	switch level {
	case DebugLevel:
		Logger.SetLevel(logrus.DebugLevel)
	case ErrorLevel:
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
}

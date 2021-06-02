package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/jonylim/basego/internal/pkg/common/constant/envvar"
)

// LogLevel is log level.
type LogLevel string

// Available log levels.
const (
	ALL   LogLevel = "ALL"
	TRACE LogLevel = "TRACE"
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
	FATAL LogLevel = "FATAL"
	OFF   LogLevel = "OFF"
)

var minLevel = 0
var mapLevels = map[LogLevel]int{
	TRACE: 0,
	DEBUG: 1,
	INFO:  2,
	WARN:  3,
	ERROR: 4,
	FATAL: 5,
}

// Init initializes the logger.
func Init() {
	// Hide log timestamp.
	HideTimestamp()

	// Set the log level.
	logLevel := LogLevel(os.Getenv(envvar.LogLevel))
	SetLogLevel(logLevel)
}

// SetLogLevel sets minimum level of logging.
func SetLogLevel(level LogLevel) {
	exists := false
	if string(level) == "" || level == ALL {
		level = ALL
		minLevel = 0
		exists = true
	} else if level == OFF {
		minLevel = 999
		exists = true
	} else {
		for l, i := range mapLevels {
			if l == level {
				minLevel = i
				exists = true
				break
			}
		}
	}
	if !exists {
		Println("logger", fmt.Sprintf(`WARN: Log level "%v" is invalid!`, level))
		level = ALL
		minLevel = 0
	}
	Println("logger", fmt.Sprintf(`Log level set to "%v".`, level))
}

// HideTimestamp hides the timestamp from printed log.
func HideTimestamp() {
	log.SetFlags(0)
}

// Println logs message without any level.
func Println(tag, message string) {
	if tag == "" {
		log.Println(message)
	} else {
		log.Println(fmt.Sprintf("[%s] %s", tag, message))
	}
}

// Trace logs message of level `TRACE`.
func Trace(tag, message string) {
	print(TRACE, tag, message)
}

// Debug logs message of level `DEBUG`.
func Debug(tag, message string) {
	print(DEBUG, tag, message)
}

// Info logs message of level `INFO`.
func Info(tag, message string) {
	print(INFO, tag, message)
}

// Warn logs message of level `WARN`.
func Warn(tag, message string) {
	print(WARN, tag, message)
}

// Error logs message of level `ERROR`.
func Error(tag, message string) {
	print(ERROR, tag, message)
}

// Fatal logs message of level `FATAL`.
func Fatal(tag, message string) {
	print(FATAL, tag, message)
}

func print(level LogLevel, tag, message string) {
	i := mapLevels[level]
	if i >= minLevel {
		if tag != "" {
			tag = " [" + tag + "]"
		}
		/* if i >= mapLevels[ERROR] {
			message = message + GenerateCallerString(3, 3)
		} */
		log.Println(string(level) + ":" + tag + " " + message)
	}
}

// FromError returns string generated from an error.
func FromError(err error) string {
	return err.Error() + GenerateCallerString(2, 3)
}

// AppendCallerString appends the file name and line number to message.
func AppendCallerString(message string) string {
	return message + GenerateCallerString(2, 3)
}

// GenerateCallerString returns string of file names and line numbers.
func GenerateCallerString(start, count int) string {
	s := ""
	for i := count; i > 0; i-- {
		_, file, line, ok := runtime.Caller(start)
		start++
		if ok {
			s = fmt.Sprintf("\n    at %v:%v", file, line) + s
		}
	}
	return s
}

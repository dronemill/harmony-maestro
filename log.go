package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

// Formatter is the log message formatter
type logFormatter struct {
}

// Format a log entry
func (c *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()
	return []byte(fmt.Sprintf("%s %s %s[%d]: %s %s\n", timestamp, hostname, logTag, os.Getpid(), strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

// logTag represents the application name generating the log message. The tag
// string will appear in all log entires.
var logTag string

func init() {
	logTag = os.Args[0]
	// log.SetFormatter(&Formatter{})
}

// LogSetTag sets the tag.
func LogSetTag(t string) {
	logTag = t
}

// LogSetLevel sets the log level. Valid levels are panic, fatal, error, warn, info and debug.
func LogSetLevel(level string) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.Fatal(fmt.Sprintf(`not a valid level: "%s"`, level))
	}
	log.SetLevel(lvl)
}

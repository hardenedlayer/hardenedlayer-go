/** file: util/debug.go
 *
 * simple logger for hardenedlayer-go
 *
 */

package util

import (
	"fmt"
	"github.com/softlayer/softlayer-go/sl"
)

const FATAL = 6
const ERROR = 5
const WARN = 4
const INFO = 3
const VERB = 2
const DEBUG = 1


type Logger struct {
	Level int
}

func GetLogger(level int) *Logger {
	debug := &Logger {
		Level: level,
	}
	return debug
}



func (d *Logger) APIError(err error) {
	e := err.(sl.Error)
	d.Error("API Error: %v (%v:%v)", e.Exception, e.StatusCode, e.Message)
}

func (d *Logger) Fatal(format string, args ...interface{}) {
	format = "FATAL: " + format
	d.Printf(FATAL, format, args...)
}

func (d *Logger) Error(format string, args ...interface{}) {
	format = "ERROR: " + format
	d.Printf(ERROR, format, args...)
}

func (d *Logger) Warn(format string, args ...interface{}) {
	format = "WARN: " + format
	d.Printf(WARN, format, args...)
}

func (d *Logger) Info(format string, args ...interface{}) {
	format = "INFO: " + format
	d.Printf(INFO, format, args...)
}

func (d *Logger) Verb(format string, args ...interface{}) {
	format = "VERB: " + format
	d.Printf(VERB, format, args...)
}

func (d *Logger) Debug(format string, args ...interface{}) {
	format = "DEBUG: " + format
	d.Printf(DEBUG, format, args...)
}

func (d *Logger) Printf(level int, format string, args ...interface{}) {
	if level >= d.Level {
		format = "--- " + format
		fmt.Printf(format, args...)
	}
}


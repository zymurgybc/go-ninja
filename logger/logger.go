package logger

import (
	"fmt"
	"os"

	"github.com/bugsnag/bugsnag-go"
	"github.com/juju/loggo"
)

// Logger wrapper for the internal logger with some extra helpers
type Logger struct {
	loggo.Logger
}

// GetLogger builds a ninja logger with the given name
func GetLogger(name string) *Logger {

	// if we are setting or configuring the ROOT logger
	if name == "" {
		// check for the magic DEBUG env variable.
		if os.Getenv("DEBUG") != "" {
			// set the default logger to debug
			loggo.GetLogger("").SetLogLevel(loggo.DEBUG)
			loggo.GetLogger("").Infof("enabling DEBUG")
		}
	}

	l := loggo.GetLogger(name)
	return &Logger{l}
}

// HandleError This notifies bugsnag and logs the error.
func (l *Logger) HandleError(err error, msg string) {
	l.Errorf("%s : %v", msg, err)
	bugsnag.Notify(err)
}

// FatalError This notifies bugsnag and logs the error then quits.
func (l *Logger) FatalError(err error, msg string) {
	l.Errorf("%s : %v", msg, err)
	bugsnag.Notify(err)
	os.Exit(1)
}

// HandleErrorf This notifies bugsnag and logs the error based on the args.
func (l *Logger) HandleErrorf(err error, msg string, args ...interface{}) {
	l.Errorf(msg, args)
	bugsnag.Notify(err)
}

// FatalErrorf This notifies bugsnag and logs the error based on the args then quits
func (l *Logger) FatalErrorf(err error, msg string, args ...interface{}) {
	l.Errorf(msg, args)
	bugsnag.Notify(err)
	os.Exit(1)
}

// FatalErrorf This notifies bugsnag and logs the error based on the args then quits
func (l *Logger) Fatalf(msg string, args ...interface{}) {
	l.Errorf(msg, args)
	bugsnag.Notify(fmt.Errorf(msg, args))
	os.Exit(1)
}

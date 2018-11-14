// Package logging implements a logging based on RFC5424 standards, simliar to
// syslog and PHP's PSR-3. This is beneficial for teams already using RFC5424 or
// PSR-3 as part of their logging practices for things such as PHP services.
// You can read more about RFC5424 at https://tools.ietf.org/html/rfc5424/ and
// PSR-3 at https://www.php-fig.org/psr/psr-3/.
package logging

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// Emergency denotes the system is unusable.
	Emergency = iota
	// Alert denotes action must be taken immediately.
	Alert
	// Critical denotes critical conditions.
	Critical
	// Error denotes error conditions.
	Error
	// Warning denotes warning conditions.
	Warning
	// Notice denotes normal but significant condition.
	Notice
	// Informational denotes informational messages.
	Informational
	// Debug denotes debug-level messages.
	Debug
)

// Level defines the severity of the message provided to the Logger.
type Level int

// Logger provides a standard set of logging functions. If you are looking to
// implement your own Logger, the StdoutLogger implementation is a good example.
type Logger interface {
	// Emergency denotes the system is unusable. Usage should accompany
	// instances that cause a restart or program exit. A Logger itself should
	// never panic or exit as part of a call to Emergency. The callee is
	// responsible to manage program exit.
	Emergency(msg string, context ...string) error
	// Alert denotes action must be taken immediately.
	Alert(msg string, context ...string) error
	// Critical denotes critical conditions.
	Critical(msg string, context ...string) error
	// Error denotes error conditions.
	Error(msg string, context ...string) error
	// Warning denotes warning conditions.
	Warning(msg string, context ...string) error
	// Notice denotes normal but significant condition.
	Notice(msg string, context ...string) error
	// Informational denotes informational messages.
	Informational(msg string, context ...string) error
	// Debug denotes debug-level messages.
	Debug(msg string, context ...string) error
	// Log provides logging based on the passed log Level. When defining a
	// custom Logger, it is recommended to implement a robust Log function and
	// then call this function within other log functions such as Error.
	Log(level Level, msg string, context ...string) error
}

// stringifyLevel is a simple helper to convert a Level to a string.
func stringifyLevel(l Level) (string, error) {
	switch l {
	case Emergency:
		return "Emergency", nil
	case Alert:
		return "Alert", nil
	case Critical:
		return "Critical", nil
	case Error:
		return "Error", nil
	case Warning:
		return "Warning", nil
	case Notice:
		return "Notice", nil
	case Informational:
		return "Informational", nil
	case Debug:
		return "Debug", nil
	default:
		return "", errors.New("Undefined log level")
	}
}

func formatMsg(level Level, msg string, context ...string) (string, error) {
	var (
		content string
		l       string
		err     error
	)
	if l, err = stringifyLevel(level); err != nil {
		return "", err
	}
	content = fmt.Sprintf("[%s] %s", l, msg)
	for _, extra := range context {
		content = strings.Join([]string{content, extra}, ", ")
	}
	return content, nil
}

type logger struct {
	Logger
	loggers []Logger
}

// NewLogger returns a singluar Logger instance that will distribute logs to all
// of the Loggers provided as arguments. If no arguments are provided, the
// StdoutLogger will be used as a default.
// Here is an example of creating a logger with multiple destinations.
//   logger := NewLogger(NewStdoutLogger(), NewPapertrailLogger(ptConfig))
func NewLogger(loggers ...Logger) Logger {
	if len(loggers) == 0 {
		return &logger{loggers: []Logger{NewStdoutLogger()}}
	} else {
		return &logger{loggers: loggers}
	}
}

func (l logger) Log(level Level, msg string, context ...string) error {
	for _, logger := range l.loggers {
		if err := logger.Log(level, msg, context...); err != nil {
			return err
		}
	}
	return nil
}

func (l logger) Emergency(msg string, context ...string) error {
	return l.Log(Emergency, msg, context...)
}
func (l logger) Alert(msg string, context ...string) error {
	return l.Log(Alert, msg, context...)
}
func (l logger) Critical(msg string, context ...string) error {
	return l.Log(Critical, msg, context...)
}
func (l logger) Error(msg string, context ...string) error {
	return l.Log(Error, msg, context...)
}
func (l logger) Warning(msg string, context ...string) error {
	return l.Log(Warning, msg, context...)
}
func (l logger) Notice(msg string, context ...string) error {
	return l.Log(Notice, msg, context...)
}
func (l logger) Informational(msg string, context ...string) error {
	return l.Log(Informational, msg, context...)
}
func (l logger) Debug(msg string, context ...string) error {
	return l.Log(Debug, msg, context...)
}

package logging

import (
	"log"
	"os"
)

// StdoutLogger is the default Logger implementation. It will take messages and
// send them to os.stdout.
type StdoutLogger struct {
	output *log.Logger
}

// NewStdoutLogger creates a new instance of the StdoutLogger.
func NewStdoutLogger() *StdoutLogger {
	return &StdoutLogger{output: log.New(os.Stdout, "", log.LstdFlags)}
}

func (logger *StdoutLogger) Log(level Level, msg string, context ...string) error {
	formattedMsg, err := formatMsg(level, msg, context...)
	if err != nil {
		return err
	}
	logger.output.Println(formattedMsg)
	return nil
}

func (logger *StdoutLogger) Emergency(msg string, context ...string) error {
	return logger.Log(Emergency, msg, context...)
}
func (logger *StdoutLogger) Alert(msg string, context ...string) error {
	return logger.Log(Alert, msg, context...)
}
func (logger *StdoutLogger) Critical(msg string, context ...string) error {
	return logger.Log(Critical, msg, context...)
}
func (logger *StdoutLogger) Error(msg string, context ...string) error {
	return logger.Log(Error, msg, context...)
}
func (logger *StdoutLogger) Warning(msg string, context ...string) error {
	return logger.Log(Warning, msg, context...)
}
func (logger *StdoutLogger) Notice(msg string, context ...string) error {
	return logger.Log(Notice, msg, context...)
}
func (logger *StdoutLogger) Informational(msg string, context ...string) error {
	return logger.Log(Informational, msg, context...)
}
func (logger *StdoutLogger) Debug(msg string, context ...string) error {
	return logger.Log(Debug, msg, context...)
}

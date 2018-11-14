package logging

// NilLogger doesn't do anything.
type NilLogger struct{}

// NewNilLogger creates a new instance of the NilLogger, which is useful for
// testing objects that require a Logger.
func NewNilLogger() *NilLogger {
	return &NilLogger{}
}

func (logger *NilLogger) Log(level Level, msg string, context ...string) error {
	return nil
}

func (logger *NilLogger) Emergency(msg string, context ...string) error {
	return logger.Log(Emergency, msg, context...)
}
func (logger *NilLogger) Alert(msg string, context ...string) error {
	return logger.Log(Alert, msg, context...)
}
func (logger *NilLogger) Critical(msg string, context ...string) error {
	return logger.Log(Critical, msg, context...)
}
func (logger *NilLogger) Error(msg string, context ...string) error {
	return logger.Log(Error, msg, context...)
}
func (logger *NilLogger) Warning(msg string, context ...string) error {
	return logger.Log(Warning, msg, context...)
}
func (logger *NilLogger) Notice(msg string, context ...string) error {
	return logger.Log(Notice, msg, context...)
}
func (logger *NilLogger) Informational(msg string, context ...string) error {
	return logger.Log(Informational, msg, context...)
}
func (logger *NilLogger) Debug(msg string, context ...string) error {
	return logger.Log(Debug, msg, context...)
}

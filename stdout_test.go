package logging

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestStdoutLogger(t *testing.T) {
	if reflect.TypeOf(NewStdoutLogger().output).Name() != reflect.TypeOf(os.Stdout).Name() {
		t.Error("wrong default output for StdoutLogger")
	}
	err := NewStdoutLogger().Log(-1, "")
	if err == nil {
		t.Error("no error returned for undefined log level")
	}
}

func TestStdoutLoggerFuncs(t *testing.T) {
	var b bytes.Buffer
	// Use a buffer instead of os.Stdout to capture output.
	bufferedLogger := &StdoutLogger{output: log.New(&b, "", 0)}
	// Define all of the required funcs for the Logger interface.
	logFuncs := []logFunc{
		logFunc(bufferedLogger.Emergency),
		logFunc(bufferedLogger.Alert),
		logFunc(bufferedLogger.Critical),
		logFunc(bufferedLogger.Error),
		logFunc(bufferedLogger.Warning),
		logFunc(bufferedLogger.Notice),
		logFunc(bufferedLogger.Informational),
		logFunc(bufferedLogger.Debug),
	}
	for _, lf := range logFuncs {
		var (
			contents []byte
			err      error
			msgs     []string
		)
		if msgs, err = callLogFunc(lf); err != nil {
			t.Error(err)
		}
		if contents, _ = ioutil.ReadAll(&b); !containsMsgs(string(contents), msgs...) {
			t.Errorf("%s was not found in %s", msgs, string(contents))
		}
	}
}

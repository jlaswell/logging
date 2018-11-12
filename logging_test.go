package logging

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type logFunc func(msg string, context ...string) error

func randString(n int) string {
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func callLogFunc(fn logFunc) ([]string, error) {
	msgs := []string{randString(6), randString(6), randString(6)}
	return msgs, fn(msgs[0], msgs[0:]...)
}

func containsMsgs(contents string, msgs ...string) bool {
	for _, msg := range msgs {
		if !strings.Contains(contents, msg) {
			return false
		}
	}
	return true
}

func captureStdout(fn func()) (captured string) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	received, _ := ioutil.ReadAll(r)
	os.Stdout = stdout
	return string(received)
}

func TestLogLevels(t *testing.T) {
	levels := map[Level]int{
		Emergency:     0,
		Alert:         1,
		Critical:      2,
		Error:         3,
		Warning:       4,
		Notice:        5,
		Informational: 6,
		Debug:         7,
	}
	for level, val := range levels {
		if int(level) != val {
			t.Errorf("level %d != %d", level, val)
		}
	}
}

func TestNewDefaultLogger(t *testing.T) {
	msg := randString(6)
	received := captureStdout(func() { NewLogger().Debug(msg) })
	expected := fmt.Sprintf("[Debug] %s", msg)
	if !strings.Contains(received, expected) {
		t.Errorf("expected %s to be contained within %s", expected, received)
	}
}

func TestUnexpectedLoggerLevel(t *testing.T) {
	if err := NewLogger().Log(-1, randString(6)); err == nil {
		t.Errorf("expected error for unexpected log level")
	}
}

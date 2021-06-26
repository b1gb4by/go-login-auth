package util

import (
	"encoding/json"
	"log"
	"os"
)

// stdLogger
// 標準出力ロガー.
type stdLogger struct {
	stderr *log.Logger
	stdout *log.Logger
}

// NewStdLogger
// コンストラクタ.
func NewStdLogger() *stdLogger {
	return &stdLogger{
		stdout: log.New(os.Stdout, "", 0),
		stderr: log.New(os.Stderr, "", 0),
	}
}

// Printf
// 標準出力 (Info).
func (l *stdLogger) Printf(format string, args ...interface{}) {
	l.stdout.Printf(format, args...)
}

// Errorf
// 標準出力 (Error).
func (l *stdLogger) Errorf(format string, args ...interface{}) {
	l.stderr.Printf(format, args...)
}

// Fatalf
// 標準出力 (Critical).
func (l *stdLogger) Fatalf(format string, args ...interface{}) {
	l.stderr.Fatalf(format, args...)
}

// AddJSONKey
// ログをjson形式で出力するために第一階層にキーを付与する.
func (l *stdLogger) AddJSONKey(key string, logVal interface{}) string {
	val := map[string]interface{}{
		key: logVal,
	}
	b, _ := json.Marshal(val)
	return string(b)
}

package log

import (
	gofmt "fmt"
	golog "log"
	"os"
	"strconv"
	"time"
)

type logger struct {
	out *golog.Logger
	err *golog.Logger
	*LoggerConfiguration
}

func NewJsonLogger(cfg *LoggerConfiguration) Logger {
	return &logger{out: golog.New(os.Stdout, "", 0), err: golog.New(os.Stderr, "", 0), LoggerConfiguration: cfg}
}

func (l *logger) SetVerbosity(Verbosity string) {
	l.Verbosity = GetVerbosityFromString(Verbosity)
}

func (l *logger) Debug(fmt string, v ...interface{}) {
	if l.Verbosity <= Debug {
		l.out.Println(l.format("debug", fmt, v...))
	}
}

func (l *logger) Info(fmt string, v ...interface{}) {
	if l.Verbosity <= Info {
		l.out.Println(l.format("info", fmt, v...))
	}
}

func (l *logger) Warn(fmt string, v ...interface{}) {
	if l.Verbosity <= Warn {
		l.out.Println(l.format("warn", fmt, v...))
	}
}

func (l *logger) Err(fmt string, v ...interface{}) {
	if l.Verbosity <= Err {
		l.err.Println(l.format("error", fmt, v...))
	}
}

func (l *logger) Panic(fmt string, v ...interface{}) {
	l.err.Panicln(l.format("panic", fmt, v...))
}

func (l *logger) Fatal(fmt string, v ...interface{}) {
	l.err.Fatalln(l.format("fatal", fmt, v...))
}

func (l *logger) format(level string, fmt string, v ...interface{}) string {
	return `{"time": "` + time.Now().Format(time.RFC3339Nano) + `", "level": "` + level + `", "message": ` + strconv.Quote(gofmt.Sprintf("%s"+fmt, append([]interface{}{l.Prefix}, v...)...)) + `}`
}

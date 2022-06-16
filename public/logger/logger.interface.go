package logger

import logrus "github.com/sirupsen/logrus"

const (
	localConfig = 1 << iota
	env
	database
	http
	all = 0xFFFFFFFF
)

var modules = map[int]string{
	localConfig: "localConfig",
	env:         "env",
	database:    "database",
	http:        "http",
}

type Logger = logrus.Logger

type ILogger interface {
	New(module int64) *Logger
	Info()
	Error()
	Trace()
	Debug()
	Warn()
	Fatal()
	Panic()
}

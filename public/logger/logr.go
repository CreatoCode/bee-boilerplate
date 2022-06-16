package logger

import "github.com/go-logr/logr"

func init() {
	logr.New().Info("logr init", "timestamp")
}

package logger

import (
	logrus "github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
	module ModuleFlag
}

var s_logger *Logger
var s_moduleFlag ModuleFlag = SharedInstance

func init() {
	logger := logrus.New()
	s_logger = &Logger{logger, SharedInstance}
	// writer1 := &bytes.Buffer{}
	// writer2 := os.Stdout
	// writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	// if err != nil {
	// 	log.Fatalf("create file log.txt failed: %v", err)
	// }

	// s_logger = log.New(io.MultiWriter(writer1, writer2, writer3), "", log.Lshortfile|log.LstdFlags)
	const msg = "logger init success"
	s_logger.Info(msg)
	s_logger.Warn(msg)
	s_logger.Trace(msg)
	s_logger.Error(msg)
}

func SetModuleEnabled(module ModuleFlag) {
	s_moduleFlag |= module
}

func New(module ModuleFlag) *Logger {
	logger := logrus.New()
	logger.SetLevel(AllLevel)
	return &Logger{logger, module}
}

func (l *Logger) Debug(args ...interface{}) {

}

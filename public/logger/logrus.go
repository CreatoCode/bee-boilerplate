package logger

import (
	logrus "github.com/sirupsen/logrus"
)

var s_logger *Logger

func init() {
	s_logger = logrus.New()
	// writer1 := &bytes.Buffer{}
	// writer2 := os.Stdout
	// writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	// if err != nil {
	// 	log.Fatalf("create file log.txt failed: %v", err)
	// }

	// s_logger = log.New(io.MultiWriter(writer1, writer2, writer3), "", log.Lshortfile|log.LstdFlags)
	s_logger.Info("logger init success")
}

func New(module int64) *Logger {
	logger := logrus.New()
	s_logger.WithFields
}

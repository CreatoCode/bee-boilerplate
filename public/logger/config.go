package logger

type config struct {
	Level      string
	Path       string
	ErrorPath  string
	AccessPath string
	MaxSize    int8 // max size of log file
	MaxAge     int8 // max age of log file
	MaxBackups int8 //
	Compress   bool // compress log file if true
}

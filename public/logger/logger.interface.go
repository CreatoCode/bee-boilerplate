package logger

type Level int

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel

	AllLevel = 0xFFFF
)

// const AllLevel = PanicLevel | FatalLevel | ErrorLevel

type ModuleFlag int64

const (
	SharedInstance ModuleFlag = 1 << iota
	LocalConfig
	Env
	Database
	Http
	All = 0xFFFFFFFF
)

var modules = map[ModuleFlag]string{
	SharedInstance: "sharedInstance",
	LocalConfig:    "localConfig",
	Env:            "env",
	Database:       "database",
	Http:           "http",
}

// type Logger interface{}

type ILogger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Trace(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

// create new logger instance
// func New(module ModuleFlag) *ILogger

// func SetModuleEnabled(module ModuleFlag)

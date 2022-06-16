package env

import (
	"bee-boilerplate/public/logger"
)

// value for the run mode in enviroment
const runEnv = "env"
const (
	dev = "dev" //  the application is in development mode
	prd = "prd" //  the application is in production mode
	stg = "stg" // the application is in staging mode
)

type IEnv interface {
	IsPrd() bool                  // IsPrd is a flag to indicate if the application is in production mode
	IsDev() bool                  // IsDev is a flag to indicate if the application is in development mode
	IsStg() bool                  // IsStg is a flag to indicate if the application is in staging mode
	CurrEnv() string              // CurrentEnviroment returns the current run mode
	Port() int                    // Port is the port number of the application
	GetStr(string, string) string // GetString returns string value of the environment variables
	GetInt(string, int) int       // GetInt returns int value of the environment variables
}

var s_logger = logger.New(logger.Env)

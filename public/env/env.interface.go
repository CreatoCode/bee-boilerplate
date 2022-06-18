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

type IEnv[T int | string] interface {
	IsPrd() bool     // IsPrd is a flag to indicate if the application is in production mode
	IsDev() bool     // IsDev is a flag to indicate if the application is in development mode
	IsStg() bool     // IsStg is a flag to indicate if the application is in staging mode
	CurrEnv() string // CurrentEnviroment returns the current run mode
	Get(T, T) T      // Get value of the environment variables
	Port() int
}

var s_logger = logger.New(logger.Env)

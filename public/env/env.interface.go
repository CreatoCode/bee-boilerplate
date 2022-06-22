package env

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
	GetInt(string, int) int       // Get value of the environment variables, T is return value if failed.
	GetStr(string, string) string // Get value of the environment variables, T is return value if failed.
	Port() int
}

package env

// type Env struct {
// 	isPrd bool // isPrd is a flag to indicate if the application is in production mode
// 	isDev bool // isDev is a flag to indicate if the application is in development mode
// 	isStg bool // isStg is a flag to indicate if the application is in staging mode
// 	Port  int  // Port is the port number of the application
// }

type IEnv interface {
	IsPrd() bool
	IsDev() bool
	IsStg() bool
	Port() int
	GetString(string, string) string // getVals returns the values of the environment variables
	GetInt(string, int) int
}

package env

import (
	"os"
	"strings"
	"sync"
)

var env sync.Map

func init() {
	for _, e := range os.Environ() {
		splits := strings.Split(e, "=")
		env.Store(splits[0], os.Getenv(splits[0]))
	}
	// env.Store("GOBIN", GetGOBIN())   // GOBIN is the path to the go binary
	// env.Store("GOPATH", GetGOPATH()) // GOPATH is the path to the go source code
}

// func getGOENV(string, error) {
// 	if file := os.Getenv("GOENV"); file != "" {
// 		if file == "off" {
// 			return "", fmt.Errorf("GOENV=off")
// 		}
// 		return file, nil
// 	}
// 	dir, err := os.UserConfigDir()
// 	if err != nil {
// 		return "", err
// 	}
// 	if dir == "" {
// 		return "", fmt.Errorf("missing user-config dir")
// 	}
// 	file, err := filepath.Join(dir, "go", "env")
// 	if err != nil {
// 		return "", err
// 	}
// 	if file == "" {
// 		return "", fmt.Errorf("missing runtime env file")
// 	}

// 	var runtimeEnv string
// 	data, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		return "", err
// 	}
// 	envStrings := strings.Split(string(data), "\n")
// 	for _, envItem := range envStrings {
// 		envItem = strings.TrimSuffix(envItem, "\r")
// 		envKeyValue := strings.Split(envItem, "=")
// 		if len(envKeyValue) == 2 && strings.TrimSpace(envKeyValue[0]) == key {
// 			runtimeEnv = strings.TrimSpace(envKeyValue[1])
// 		}
// 	}
// 	return runtimeEnv, nil
// }

func Get(key string, defaultValue string) string {
	if val, ok := env.Load(key); ok {
		return val.(string)
	}
	return defaultValue
}

// func GetGOBIN() string {
// 	// The one set by user explicitly by `export GOBIN=/path` or `env GOBIN=/path command`
// 	gobin := strings.TrimSpace(Get("GOBIN", ""))
// 	if gobin == "" {
// 		var err error
// 		// The one set by user by running `go env -w GOBIN=/path`
// 		gobin, err = GetRuntimeEnv("GOBIN")
// 		if err != nil {
// 			// The default one that Golang uses
// 			return filepath.Join(build.Default.GOPATH, "bin")
// 		}
// 		if gobin == "" {
// 			return filepath.Join(build.Default.GOPATH, "bin")
// 		}
// 		return gobin
// 	}
// 	return gobin
// }

// // GetGOPATH returns GOPATH environment variable as a string.
// // It will NOT be an empty string.
// func GetGOPATH() string {
// 	// The one set by user explicitly by `export GOPATH=/path` or `env GOPATH=/path command`
// 	gopath := strings.TrimSpace(Get("GOPATH", ""))
// 	if gopath == "" {
// 		var err error
// 		// The one set by user by running `go env -w GOPATH=/path`
// 		gopath, err = GetRuntimeEnv("GOPATH")
// 		if err != nil {
// 			// The default one that Golang uses
// 			return build.Default.GOPATH
// 		}
// 		if gopath == "" {
// 			return build.Default.GOPATH
// 		}
// 		return gopath
// 	}
// 	return gopath
// }

func Port() int {
	return GetInt("port", 8080)
}

func GetString(key string, defaultValue string) string {
	if val, ok := env.Load(key); ok {
		return val.(string)
	}
	return defaultValue
}

func GetInt(key string, defaultValue int) int {
	if val, ok := env.Load(key); ok {
		return val.(int)
	}
	return defaultValue
}

func IsPrd() bool {
	return GetString(runEnv, prd) == prd
}

func IsDev() bool {
	return GetString(runEnv, prd) == dev
}

func IsStg() bool {
	return GetString(runEnv, prd) == stg
}

func CurrEnv() string {
	return GetString(runEnv, prd)
}

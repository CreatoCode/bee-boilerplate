package env

import (
	"bee-boilerplate/public/logger"
	"os"
	"strconv"
	"strings"
	"sync"
)

type env struct {
	syncMap sync.Map
	IEnv
}

var g_env = &env{}

var g_logger = logger.New(logger.Env)

func init() {
	for _, e := range os.Environ() {
		splits := strings.Split(e, "=")
		g_env.syncMap.Store(splits[0], os.Getenv(splits[0]))
	}
	g_logger.Debug(os.Environ())
	g_logger.Debug("init done")
	// env.Store("GOBIN", GetGOBIN())   // GOBIN is the path to the go binary
	// env.Store("GOPATH", GetGOPATH()) // GOPATH is the path to the go source code
}

func SharedInstance() IEnv {
	return g_env
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

// func GetGOBIN() string {
// 	// The one set by user explicitly by `export GOBIN=/path` or `env GOBIN=/path command`
// 	gobin := strings.TrimSpace(get("GOBIN", ""))
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
// 	gopath := strings.TrimSpace(get("GOPATH", ""))
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

func get[T int | string](key string, defaultValue T) T {
	if val, ok := g_env.syncMap.Load(key); ok {
		switch any(defaultValue).(type) {
		case int:
			if i, err := strconv.Atoi(val.(string)); err == nil {
				return T(i)
			}

		}
		return val.(T)
	}
	return defaultValue
}

func (e *env) IsPrd() bool {
	return get(runEnv, prd) == prd
}

func (e *env) IsDev() bool {
	return get(runEnv, prd) == dev
}

func (e *env) IsStg() bool {
	return get(runEnv, prd) == stg
}

func (e *env) CurrEnv() string {
	return get(runEnv, prd)
}

func (e *env) Port() int {
	return get("port", 8080)
}

func (e *env) GetInt(key string, defaultValue int) int {
	return get(key, defaultValue)
}

func (e *env) GetStr(key string, defaultValue string) string {
	return get(key, defaultValue)
}

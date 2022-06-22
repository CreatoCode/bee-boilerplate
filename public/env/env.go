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
}

var g_logger = logger.New(logger.Env)

var g_env = &env{}

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

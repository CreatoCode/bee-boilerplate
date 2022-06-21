package ORM

import (
	"bee-boilerplate/public/errors"
	"bee-boilerplate/public/logger"
)

var g_logger = logger.New(logger.ORM)

var g_createError = errors.NewDomain("bee-boilerplate/public/datalayer/orm")

const (
	ErrorCodeOK = iota + 1
)

type orm[T any] struct {
	orm *T
}

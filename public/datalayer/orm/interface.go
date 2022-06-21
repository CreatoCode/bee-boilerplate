package ORM

import (
	"bee-boilerplate/public/error"
	"bee-boilerplate/public/logger"
)

var g_logger = logger.New(logger.ORM)

var g_createError = error.NewDomain("bee-boilerplate.public.orm")

const (
	ErrorCodeOK = iota + 1
)

// ORM abstract

type IORM[T any] interface {
	Create(value interface{}) *error.Error
}

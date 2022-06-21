// data layer abstract
package datalayer

import "bee-boilerplate/public/logger"

type IDatalayer[T any] interface {
	Create() error
	Get() (T, error)
	Update(T) error
	Delete(int64) error
}

var g_logger = logger.New(logger.DataLayer)

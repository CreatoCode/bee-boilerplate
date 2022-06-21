// data layer abstract
package datalayer

import (
	"bee-boilerplate/public/errors"
	"bee-boilerplate/public/logger"
)

type IDatalayer interface {
	Create(value interface{}) *errors.Error
	Get(model interface{}, conds ...interface{}) ([]interface{}, *errors.Error)
	GetFirst(model interface{}, conds ...interface{}) (interface{}, *errors.Error)
	Update(column string, model interface{}) *errors.Error
	Delete(model interface{}, conds ...interface{}) *errors.Error
}

var g_logger = logger.New(logger.DataLayer)
var g_errors = errors.NewDomain("bee-boilerplate/public/datalayer")

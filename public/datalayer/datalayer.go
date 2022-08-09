package datalayer

import (
	"bee-boilerplate/public/datalayer/orm/gorm"
	"bee-boilerplate/public/errors"
	"bee-boilerplate/public/logger"
)

var g_logger = logger.New(logger.DataLayer)
var g_errors = errors.NewDomain("bee-boilerplate/public/datalayer")

var g_orm = gorm.SharedInstance()

func Create(value interface{}) *errors.Error {
	return g_orm.Create(value)
}

func Get(model interface{}, conds ...interface{}) ([]interface{}, *errors.Error) {
	return g_orm.Get(model, conds...)
}

func GetFirst(model interface{}, conds ...interface{}) (interface{}, *errors.Error) {
	return g_orm.GetFirst(model, conds...)
}

func Update(column string, model interface{}) *errors.Error {
	return g_orm.Update(column, model)
}

func Delete(model interface{}, conds ...interface{}) *errors.Error {
	return g_orm.Delete(model, conds...)
}

func AutoMigrate(value interface{}) *errors.Error {
	return g_orm.AutoMigrate(value)
}

func IsExist(model interface{}, conds ...interface{}) (bool, *errors.Error) {
	return g_orm.IsExist(model, conds...)
}

package datalayer

import (
	ORM "bee-boilerplate/public/datalayer/orm"
	"bee-boilerplate/public/errors"
)

type datalayer[T any] struct {
	orm *T
}

var g_orm = ORM.ShareInstance()

var g_datalayer = &datalayer[any]{orm: g_orm}

func shareInstance() *datalayer {
	return g_datalayer
}

func (d *datalayer[T]) Create(value interface{}) *errors.Error {
	return g_orm.Create(value)
}

func (d *datalayer[T]) Get(model interface{}, conds ...interface{}) ([]interface{}, *errors.Error) {
	return g_orm.Get(model, conds...)
}

func (d *datalayer[T]) GetFirst(model interface{}, conds ...interface{}) (interface{}, *errors.Error) {
	return g_orm.GetFirst(model, conds...)
}

func (d *datalayer[T]) Update(column string, model interface{}) *errors.Error {
	return g_orm.Update(column, model)
}

func (d *datalayer[T]) Delete(model interface{}, conds ...interface{}) *errors.Error {
	return g_orm.Delete(model, conds...)
}

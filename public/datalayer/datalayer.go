package datalayer

import (
	ORM "bee-boilerplate/public/datalayer/orm"
	"bee-boilerplate/public/errors"
)

type datalayer struct {
}

var g_orm = ORM.New()

func New() *datalayer {
	return &datalayer{}
}

func (d *datalayer) Create(value interface{}) *errors.Error {
	return g_orm.Create(value)
}

func (d *datalayer) Get(model interface{}, conds ...interface{}) ([]interface{}, *errors.Error) {
	return g_orm.Get(model, conds...)
}

func (d *datalayer) GetFirst(model interface{}, conds ...interface{}) (interface{}, *errors.Error) {
	return g_orm.GetFirst(model, conds...)
}

func (d *datalayer) Update(column string, model interface{}) *errors.Error {
	return g_orm.Update(column, model)
}

func (d *datalayer) Delete(model interface{}, conds ...interface{}) *errors.Error {
	return g_orm.Delete(model, conds...)
}

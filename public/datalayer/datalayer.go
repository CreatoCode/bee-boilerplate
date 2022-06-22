package datalayer

import (
	"bee-boilerplate/public/datalayer/interface/iface"
	"bee-boilerplate/public/datalayer/orm/gorm"
	"bee-boilerplate/public/errors"
	"bee-boilerplate/public/logger"
)

var g_logger = logger.New(logger.DataLayer)
var g_errors = errors.NewDomain("bee-boilerplate/public/datalayer")

type datalayer[T iface.Model] struct {
	model T
	orm   iface.IDatalayer[T]
}

func New[T iface.Model](model iface.Model) iface.IDatalayer[iface.Model] {
	return &datalayer[model]{model: model, orm: gorm.New(model)}
}

func (d *datalayer[T]) Create(model iface.Model) *errors.Error {
	return datalayer[iface.Model].orm.Create(model)
}

func (d *datalayer[T]) Get(model iface.Model, conds ...interface{}) ([]iface.Model, *errors.Error) {
	results, err := datalayer[iface.Model].orm.Get(model, conds...)
	s := make([]iface.Model, 1)
	return s, nil
}

func (d *datalayer[T]) GetFirst(model iface.Model, conds ...interface{}) (iface.Model, *errors.Error) {
	return datalayer[iface.Model].orm.GetFirst(model, conds...)
}

func (d *datalayer[T]) Update(column string, model iface.Model) *errors.Error {
	return datalayer[iface.Model].orm.Update(column, model)
}

func (d *datalayer[T]) Delete(model iface.Model, conds ...interface{}) *errors.Error {
	return datalayer[iface.Model].orm.Delete(model, conds...)
}

// func (d *datalayer) Create(value interface{}) *errors.Error {
// 	return g_orm.Create(value)
// }

// func (d *datalayer) Get(model interface{}, conds ...interface{}) ([]interface{}, *errors.Error) {
// 	return g_orm.Get(model, conds...)
// }

// func (d *datalayer) GetFirst(model interface{}, conds ...interface{}) (interface{}, *errors.Error) {
// 	return g_orm.GetFirst(model, conds...)
// }

// func (d *datalayer) Update(column string, model interface{}) *errors.Error {
// 	return g_orm.Update(column, model)
// }

// func (d *datalayer) Delete(model interface{}, conds ...interface{}) *errors.Error {
// 	return g_orm.Delete(model, conds...)
// }

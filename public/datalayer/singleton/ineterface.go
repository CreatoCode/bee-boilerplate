// data layer abstract
package singleton

import (
	"bee-boilerplate/public/errors"
)

type IDatalayer interface {
	AutoMigrate(value ...interface{}) *errors.Error
	Create(value interface{}) *errors.Error
	Get(model interface{}, conds ...interface{}) ([]interface{}, *errors.Error)
	GetFirst(model interface{}, conds ...interface{}) (interface{}, *errors.Error)
	Update(column string, model interface{}) *errors.Error
	Delete(model interface{}, conds ...interface{}) *errors.Error
}

// type Model any

// type IDatalayer[T Model] interface {
// 	Create(model T) *errors.Error
// 	Get(model T, conds ...interface{}) ([]T, *errors.Error)
// 	GetFirst(model T, conds ...interface{}) (T, *errors.Error)
// 	Update(column string, model T) *errors.Error
// 	Delete(model T, conds ...interface{}) *errors.Error
// }

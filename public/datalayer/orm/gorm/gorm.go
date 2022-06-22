package gorm

import (
	"bee-boilerplate/public/datalayer/interface/iface"
	ORM "bee-boilerplate/public/datalayer/orm"
	"bee-boilerplate/public/errors"
	"bee-boilerplate/public/logger"

	"gorm.io/driver/sqlite"
	goorm "gorm.io/gorm"
)

type orm[T iface.Model] struct {
	model T
}

var g_createError = errors.NewDomain("bee-boilerplate/public/datalayer/orm")
var g_logger = logger.New(logger.ORM)
var g_gorm *goorm.DB

func init() {
	// g_db, err = gorm.Open("mysql", "root:bgbiao.top@(127.0.0.1:13306)/test_api?charset=utf8&parseTime=True&loc=Local")
	db, err := goorm.Open(sqlite.Open("test.db"), &goorm.Config{})
	if err != nil {
		g_logger.Error("创建数据库连接失败:%v", err)
	} else {
		g_gorm = db
	}
}

func New[T iface.Model](model iface.Model) iface.IDatalayer {
	return &orm{model: model}
}

func (o *orm[T]) Create(value iface.Model) *errors.Error {
	db := g_gorm.Create(value)
	if nil != db.Error {
		err := g_createError(ORM.ErrorCodeOK, "创建失败", "", "")
		return &err
	}
	return nil
}

func (o *orm[T]) Get(model iface.Model, conds ...interface{}) ([]iface.Model, *errors.Error) {
	db := g_gorm.Take(model, conds...)
	if nil != db.Error {
		err := g_createError(ORM.ErrorCodeOK, "查询失败", "", "")
		return nil, &err
	}
	return nil, nil
}

func (d *orm[T]) GetFirst(model iface.Model, conds ...interface{}) (iface.Model, *errors.Error) {
	db := g_gorm.First(model, conds...)
	if nil != db.Error {
		err := g_createError(ORM.ErrorCodeOK, "查询失败", "", "")
		return nil, &err
	}
	return nil, nil
}

func (d *orm[T]) Update(column string, model iface.Model) *errors.Error {
	db := g_gorm.Update(column, model)
	if nil != db.Error {
		err := g_createError(ORM.ErrorCodeOK, "更新失败", "", "")
		return &err
	}
	return nil
}

func (d *orm[T]) Delete(model iface.Model, conds ...interface{}) *errors.Error {
	db := g_gorm.Delete(model, conds...)
	if nil != db.Error {
		err := g_createError(ORM.ErrorCodeOK, "删除失败", "", "")
		return &err
	}
	return nil
}

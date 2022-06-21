package ORM

import (
	"bee-boilerplate/public/errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type orm struct {
	orm *gorm.DB
}

var g_db *orm

func init() {
	// g_db, err = gorm.Open("mysql", "root:bgbiao.top@(127.0.0.1:13306)/test_api?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		g_logger.Error("创建数据库连接失败:%v", err)
	} else {
		g_db = &orm{orm: db}
	}
}

func New() *orm {
	return g_db
}

func (o *orm) Create(value interface{}) *errors.Error {
	db := g_db.orm.Create(value)
	if nil != db.Error {
		err := g_createError(ErrorCodeOK, "创建失败", "", "")
		return &err
	}
	return nil
}

func (o *orm) Get(model interface{}, conds ...interface{}) ([]interface{}, *errors.Error) {
	db := g_db.orm.Take(model, conds...)
	if nil != db.Error {
		err := g_createError(ErrorCodeOK, "查询失败", "", "")
		return nil, &err
	}
	return nil, nil
}

func (d *orm) GetFirst(model interface{}, conds ...interface{}) (interface{}, *errors.Error) {
	db := g_db.orm.First(model, conds...)
	if nil != db.Error {
		err := g_createError(ErrorCodeOK, "查询失败", "", "")
		return nil, &err
	}
	return nil, nil
}

func (d *orm) Update(column string, model interface{}) *errors.Error {
	db := g_db.orm.Update(column, model)
	if nil != db.Error {
		err := g_createError(ErrorCodeOK, "更新失败", "", "")
		return &err
	}
	return nil
}

func (d *orm) Delete(model interface{}, conds ...interface{}) *errors.Error {
	db := g_db.orm.Delete(model, conds...)
	if nil != db.Error {
		err := g_createError(ErrorCodeOK, "删除失败", "", "")
		return &err
	}
	return nil
}

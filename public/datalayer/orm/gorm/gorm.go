package gorm

import (
	ORM "bee-boilerplate/public/datalayer/orm"
	"bee-boilerplate/public/datalayer/singleton"
	"bee-boilerplate/public/errors"
	"bee-boilerplate/public/logger"

	"gorm.io/driver/sqlite"
	goorm "gorm.io/gorm"
)

var g_createError = errors.NewDomain("bee-boilerplate/public/datalayer/gorm")
var g_logger = logger.New(logger.ORM)

type orm struct {
	orm *goorm.DB
}

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

func SharedInstance() singleton.IDatalayer {
	return &orm{orm: g_gorm}
}

func (o *orm) Create(value interface{}) *errors.Error {
	db := g_gorm.Create(value)
	if nil != db.Error {
		err := g_createError(ORM.ErrorCodeOK, "创建失败", "", "")
		return &err
	}
	return nil
}

func (o *orm) Get(model interface{}, conds ...interface{}) ([]interface{}, *errors.Error) {
	db := g_gorm.Take(model, conds...)
	if nil != db.Error {
		err := g_createError(ORM.ErrorCodeOK, "查询失败", "", "")
		return nil, &err
	}
	return nil, nil
}

func (d *orm) GetFirst(model interface{}, conds ...interface{}) (interface{}, *errors.Error) {
	db := g_gorm.First(model, conds...)
	if nil != db.Error {
		err := g_createError(ORM.ErrorCodeOK, "查询失败", "", "")
		return nil, &err
	}
	return nil, nil
}

func (d *orm) Update(column string, model interface{}) *errors.Error {
	db := g_gorm.Update(column, model)
	if nil != db.Error {
		err := g_createError(ORM.ErrorCodeOK, "更新失败", "", "")
		return &err
	}
	return nil
}

func (d *orm) Delete(model interface{}, conds ...interface{}) *errors.Error {
	db := g_gorm.Delete(model, conds...)
	if nil != db.Error {
		err := g_createError(ORM.ErrorCodeOK, "删除失败", "", "")
		return &err
	}
	return nil
}

func (d *orm) AutoMigrate(value ...interface{}) *errors.Error {
	err := g_gorm.AutoMigrate(value...)
	if nil != err {
		err := g_createError(ORM.ErrorCodeOK, "删除失败", err.Error(), "")
		return &err
	}
	return nil
}

func (d *orm) IsExist(model interface{}, conds ...interface{}) (bool, *errors.Error) {
	// g_gorm.Raw("SELECT EXISTS(SELECT 1 FROM my_table WHERE id = ? AND `key` = ? AND `value` = ?) AS found",
	// myID, myKey, "0").Scan(&result)
	// var result struct {
	// 	Found bool
	// }
	// g_gorm.Raw("SELECT 1 FROM my_table WHERE id = ? AND `key` = ? AND `value` = ?",
	// 	myID, myKey, "0").Scan(&result)
	countryCapitalMap := make(map[string]string)
	err := g_gorm.Model(model).Where(conds[0], conds[1:]...).Select("1").Scan(&countryCapitalMap).Error
	if err != nil {
		err := g_createError(ORM.ErrorCodeOK, "查询失败", err.Error(), "")
		return false, &err
	}
	return true, nil
}

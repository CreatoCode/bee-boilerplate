package ORM

import (
	err "bee-boilerplate/public/error"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var g_db *gorm.DB

func init() {
	var err error
	// g_db, err = gorm.Open("mysql", "root:bgbiao.top@(127.0.0.1:13306)/test_api?charset=utf8&parseTime=True&loc=Local")
	g_db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		g_logger.Error("创建数据库连接失败:%v", err)
	}
}

func Create(value interface{}) *err.Error {
	db := g_db.Create(value)
	if nil != db.Error {
		err := g_createError(ErrorCodeOK, "创建失败", "", "")
		return &err
	}
	return nil
}

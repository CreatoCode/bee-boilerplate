package ORM

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:bgbiao.top@(127.0.0.1:13306)/test_api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		s_logger.Error("创建数据库连接失败:%v", err)
	}
	db.Create()
}

package main // 声明 main 包，表明当前是一个可执行程序

import (
	"bee-boilerplate/public/datalayer"
	"bee-boilerplate/public/env" // 导入内置 fmt 包
	http "bee-boilerplate/public/net/http/server"
	"bee-boilerplate/service/user"
	"fmt"
)

type Test struct {
	Name string
}

func registerRouter() {
	engine := http.Default()

	user.RegisterRouter(engine.Group("/user"))

	engine.Run(":9999")
}

func main() { // main函数，是程序执行的入口
	t := Test{Name: "baby"}
	// datalayer := datalayer.New[Test](t)
	datalayer.AutoMigrate(&t)
	datalayer.Create(t)
	env := env.SharedInstance()
	fmt.Println("Hello World! ", env.Port(), env.CurrEnv(), env.IsDev(), env.IsPrd(), env.IsStg()) // 在终端打印 Hello World!
	registerRouter()
}

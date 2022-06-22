package main // 声明 main 包，表明当前是一个可执行程序

import (
	"bee-boilerplate/public/datalayer"
	"bee-boilerplate/public/env" // 导入内置 fmt 包
	"fmt"
)

type Test struct {
	Name string
}

func main() { // main函数，是程序执行的入口
	t := Test{Name: "baby"}
	datalayer := datalayer.New[Test](t)
	datalayer.Create(t)
	env := env.SharedInstance()
	fmt.Println("Hello World! ", env.Port(), env.CurrEnv(), env.IsDev(), env.IsPrd(), env.IsStg()) // 在终端打印 Hello World!
}

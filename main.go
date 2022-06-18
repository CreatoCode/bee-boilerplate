package main // 声明 main 包，表明当前是一个可执行程序

import (
	"bee-boilerplate/public/env" // 导入内置 fmt 包
	"fmt"
)

func main() { // main函数，是程序执行的入口
	fmt.Println("Hello World! ", env.Port(), env.CurrEnv(), env.IsDev(), env.IsPrd(), env.IsStg()) // 在终端打印 Hello World!
}

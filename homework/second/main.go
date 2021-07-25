package main

import (
	"fmt"
	"go_lessons/homework/second/dao"
)

func main() {
	//调用功能函数
	err := dao.TestFunction()
	//打印错误堆栈信息
	if err != nil {
		fmt.Printf("\n%+v\n", err)
	}
}

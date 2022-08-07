package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("./dll.so")
	if err != nil {
		panic(err)
	}

	// 1 查找方法
	target, err := p.Lookup("Greet")
	if err != nil {
		panic(err)
	}

	fn, ok := target.(func(int) string)
	if !ok {
		panic("symbol not consistent")
	}
	fmt.Printf("函数类型: %T, 调用结果: %s\n", target, fn(1024))

	// 2 查找变量
	// 2.1 查找小写字母开头命名的变量
	internalVal, err := p.Lookup("internalVal")
	if err != nil {
		fmt.Printf("小写字母开头命名的变量查找失败: %s\n", err)
	} else {
		fmt.Printf("小写字母开头命名的变量: %v\n", internalVal)
	}

	// 2.2 查找大写字母开头命名的变量
	ExportVal, err := p.Lookup("ExportVal")
	if err != nil {
		fmt.Printf("大写字母开头命名的变量查找失败: %s\n", err)
	} else {
		fmt.Printf("大写字母开头命名的变量: %T %d\n", ExportVal, *(ExportVal.(*int)))
	}

	// output:
	// 函数类型: func(int) string, 调用结果: hello No.1024
	// 小写字母开头命名的变量查找失败: plugin: symbol internalVal not found in plugin plugin/unnamed-cafeff5bb7c82481370af4eb8b9407a88f518582
	// 大写字母开头命名的变量: *int 12
}

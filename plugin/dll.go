package main

import (
	"strconv"
)

// 动态库接口，使用如下命令生成.so文件：go build -buildmode=plugin dll.go
// Note:
// 1. 该plugin必须有main包
// 2. 符号不支持方法，只能是普通函数

var (
	internalVal = 11 // 不可被导入
	ExportVal   = 12 // 导入后类型为*int，获取值时需要解引用
)

func Greet(no int) string {
	return "hello No." + strconv.Itoa(no)
}

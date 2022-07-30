package go_advance

import (
	"fmt"
)

var (
	version   string
	user      string
	commitId  string
	buildTime string
)

const level int = 1

// 通过go build命令将动态数据注入到二进制文件中，其中-ldflags会将参数传递给连接器linker，
// 可以参考这篇文章：https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
// -ldflag其他选项，请参考：https://pkg.go.dev/cmd/link 或者在命令行中使用：go tool link --help
// 下面是go build命令格式：
// go build -ldflags="-X 'main.version=$(git describe --tags --abbrev=0)' \
// -X 'main.commitId=$(git rev-parse HEAD)' -X 'main.user=$(id -u -n)' \
// -X 'main.buildTime=$(date)' -X 'main.level=3'"
//
// 注意：
// 1. 常量的值不会被替换，不过go build命令中使用-ldflags修改常量的值并不会报错
// 2. 如果需要替换子包中的变量的话，格式如下-X 'app/build.foo=anything'
// 3. -ldflags中-X只能用来传递字符串，不能传递int等类型数据
func buildInfo() {
	fmt.Printf("%13s%s\n", "version: ", version)
	fmt.Printf("%13s%s\n", "commit id: ", commitId)
	fmt.Printf("%13s%d\n", "user: ", user)
	fmt.Printf("%13s%s\n", "build time: ", buildTime)
	fmt.Printf("%13s%d\n", "const value: ", level)

	// maybe output:
	//     version: v1.1.0
	//   commit id: 1481a0b46d26d26deced4b99c072bf33e53d20d7
	//        user: qshuai
	//  build time: 2022年 7月30日 星期六 14时12分19秒 CST
	// const value: 1
}

dll.go文件为生成动态链接库的源码文件，通过如下命令生成.so文件: 
```shell
go build -buildmode=plugin dll.go
```

main.go为导入动态库的主程序源码文件，通过运行命令`go run main.go`来查看运行结果，辅助inspect动态链接库的使用

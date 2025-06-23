package main
import "fmt"
// 单元测试 go test
// 使用该命令，所有以_test.go为后缀的源码文件都会被go test运行到
//  go build命令不会将这些测试文件打包到最后的可执行文件中
//test文件有4类
// Test开头   功能测试； Benchmark开头 性能测试；example开头 模糊测试；
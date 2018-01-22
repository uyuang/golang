//当前程序的包名
package main

//导入的其他包
import (
	"fmt"
)

//import "os"
//import "time"
//import "strings"
//常量的定义
const (
	pi    = 3.14
	testa = 3.15
)

//全局变量
var (
	name     = "gophe111r"
	testname = "testname"
)

//一般类型声明
type (
	newType  int
	testType int
	newA     int8
)

//结构声明
type (
	gopher struct{}
	aa     struct{}
)

//接口声明
type (
	golang  interface{}
	golanga interface{}
)

func main() {
	fmt.Println("hello world!你好，世界！")

}

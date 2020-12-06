//要求开发一个hello.go 程序
//package main 表示 hello.go 文件所在的包是main，在go中每个文件必须归属一个包
package main

import (
	ro "Learngolang/src/go_code/goTutorialProject/header"
	"fmt"
)

func main()  {
	ro.Print("hello world go!!!!")
	ro.PrintInput()
	fmt.Println()
}

//使用 go build hello.dgo 进行编译
//使用自定义名字 go build -o myhello.exe hello.go 
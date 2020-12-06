package main

import "fmt"

func main()  {
	//默认变量10
	var a int = 10
	//声明一个空指针
	var ptr *int
	//指针取a的地址
	ptr = &a
	//修改 a 的内存地址 的变量为30
	*ptr = 30
	fmt.Println(a) //--》 30
	fmt.Println(*ptr) // --》30
}
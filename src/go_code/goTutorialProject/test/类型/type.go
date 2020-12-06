package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var c = 100
	//输出类型
	fmt.Printf(" c 的类型是 %T \n ",c)
	//------------------
	//如何在程序查看某个变量的占用字节大小和数据类型
	var c2 int64 = 10
	//unsafe.Sizeof(c2) 是unsafe包的一个函数，可以返回c2变量占用的字节数
	fmt.Printf("c 的类型 %T \n c2 占用的字节数是 %d ", c , unsafe.Sizeof(c2))

	//实际开发中，尽量使用占用空间小的数据类型
	var age byte = 90  // byte 0~255  //bit是最小的单位   1 byte = 8 bit
	fmt.Println(age)
}

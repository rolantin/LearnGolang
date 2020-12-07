package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var a bool = false
	//注意 bool类型占用存储空间是1个字节
	fmt.Println("a 的占用空间= " ,unsafe.Sizeof(a))
    //bool 只能取false true
}
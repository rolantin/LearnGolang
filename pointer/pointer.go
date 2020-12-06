package main

import "fmt"

func addOne(num int) {
	num=num+1
}

func addOneWithPointer(num *int) {
	*num=*num+1
}

func main()  {
	x:=5
	fmt.Println(x)

	//xPrt:= &x
	var xPrt *int = &x  //因为指针源为int 所以 *int 这里打印的是内存地址
	fmt.Println(xPrt)
	addOne(x)
	fmt.Println(x , "值为5")
	//可修改内存的值进行加法，使其生效
	addOneWithPointer(xPrt)
	fmt.Println(x ,"值为6")
}

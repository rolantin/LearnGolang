package main

import "fmt"

func  main()  {
	//a := 200
	//b := &a
	//fmt.Println("a -->", a)
	//fmt.Println("*b -->", *b) //读取内存&a 值
	//fmt.Println("&a -->", &a)
	//fmt.Println("b -->", b) //地址
	//*b++
	//fmt.Println("a -->", a)
	//fmt.Println("*b -->", *b)

	c:= 300
	b:= &c
	fmt.Println(c,*b)
	*b = 500
	fmt.Println(c,*b)
	fmt.Println(&c)
}

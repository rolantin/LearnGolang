package main

import "fmt"

//1.对于基本数据类型，变量存的就是值，也叫值类型
//2.获取变量的地址，用& ，比如 var num int，获取num 的地址：&num
//3.指针变量存的是一个地址，这个地址指向的空间存的才是值 比如 var ptr *int = &num
//4.获取指针类型所指的值，使用： * 比如: var ptr *int ,使用 *ptr 获取ptr指向的值

//演示golang中指针类型
func main()  {
	//基本数据类型在内存布局
	var i int = 10
	// i 的地址是什么，&i
	fmt.Println(" i的地址= ",&i)
    // 下面的 var ptr *int = &i
    //1.ptr 是一个指针变量
    //2.ptr 的类型 *int
    //3.ptr 本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr=%v",ptr)
	fmt.Printf("ptr的地址=%v",&ptr)
	fmt.Printf("ptr指向的值=%v",*ptr)


}
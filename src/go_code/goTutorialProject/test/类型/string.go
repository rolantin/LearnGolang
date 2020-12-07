package main

import "fmt"

func main()  {
	//go中一旦定义了字符串就不能改变了
	var c string = "aaa/naaa"
	fmt.Println(c)
	//str[0] = 'a' 这里就不能去修改str的内容，即go中的字符串不可变
	//输出源代码效果 使用 反引号
	str2:= `package main
	import "fmt"
	func main()  {
		//go中一旦定义了字符串就不能改变了
		var c string = "aaa/naaa"
		fmt.Println(c)
		//str[0] = 'a' 这里就不能去修改str的内容，即go中的字符串不可变
		//输出源代码效果 使用 反引号'`

	//注意 字符串拼接方式
	/*
	str:= "hellp" + "world"
	正确的
	str:= "hellp" +
		     "world"
	错误的
	str:= "hellp"
		     +"world" */
}
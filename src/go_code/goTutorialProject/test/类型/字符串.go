package main

import "fmt"

func main(){
	var c1 byte = 'a'
	var c2 byte = '0'
	//当我们输出byte值 就是输出了对应的字符的码值
	//ASCII表可查
	fmt.Println("c1=",c1) //--》 97
	fmt.Println("c2=",c2) //--》 48
	//如果希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c",c1,c2) //--> a 0

	//如果中文
	//var c3 byte = '北' //--> 提示溢出
	var c4 int= '北'
	fmt.Printf("c1=%c c3对应码=%d",c4,c4) //--> a 0
	//在ascii 表 0-1 a-z A-Z 可以直接用byte
	//如果我们保存的字节码大于255值，可以考虑用int 输出字节符
    //我们需要输出字符 使用 %c
}
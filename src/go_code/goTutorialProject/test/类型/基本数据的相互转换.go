package main

import (
	"fmt"
)

func main()  {
	var i int32 = 100
	var nl float32 = float32(i)
	var nl2 int8 = int8(nl)
    //指令格式
	fmt.Println("i=%v nl=%v nl2=%v",i,nl,nl2)

	//nl = strconv.FormatInt(int64(i),10)

}
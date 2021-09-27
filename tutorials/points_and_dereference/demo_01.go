package main

import "fmt"

func changeValue(str *string) {
	*str = "changed"
}

func changeValue1(str string) string {
	str = "changed one"
	return str
}

func main() {
	//
	// 将数值 7 赋值给变量
	var x uint8 = 7
	// & 将变量 x 的内存地址，取址符号
	var y *uint8 = &x
	// fmt.Println(&x) //0xc000014090

	// *y = 8

	fmt.Println(x, y)

	toChange := "hello"
	changeValue(&toChange)

	fmt.Println(toChange)

	changeValue1(toChange)
	fmt.Println(toChange)
}

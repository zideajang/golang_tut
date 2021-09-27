package main

import "fmt"

func test(){
	fmt.Println("test function")
}

func testTwo(myFun func (int)int) int{ 
	return myFun(2)	
}

func main()  {
	x := test
	x()
	a := func ()  {
		fmt.Println("a function")
	}
	a()

	res := func (x int) int {
		return x * 2
	}

	resOne := testTwo(res)

	fmt.Println(resOne)
}

//ssh-keygen -t rsa
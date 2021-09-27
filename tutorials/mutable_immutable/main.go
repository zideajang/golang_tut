package main

import "fmt"

func main() {
	// var x int = 2
	// y := x
	// y = 100

	// var x []int = []int{2,3,5}
	// y := x
	// y[0] = 100

	var x map[string]int = map[string]int{"hello": 123}
	y := x
	y["machine"] = 2
	// x["learning"] = 3

	fmt.Println(x, y)
}

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 45, 6, 7, 8, 9}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, element := range arr {
		fmt.Println(arr[i])
		fmt.Println(element)
	}
}

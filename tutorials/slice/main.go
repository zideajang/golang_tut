package main

import "fmt"

func appendSlice() []int {
	a := []int{1, 2, 3, 4, 5}
	b := []int{7, 8, 9}
	return append(a, b...)
}

func copySliceV1() []int {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func copySliceV2() []int {
	a := []int{1, 2, 3, 4, 5}
	return append(a[:0:0], a...)
}

func copySliceV3() []int {
	a := []int{1, 2, 3, 4, 5}
	return append([]int(nil), a...)
}

func outSlice() []int {
	a := []int{1, 2, 3, 4, 5}
	i := 1
	j := 4

	return append(a[:i], a[j:]...)

}

func main() {

	s2 := []int{0, 1, 2, 5: 100}
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 3, 3)
	s5 := make([]int, 3, 6)

	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s5, len(s5), cap(s5))

	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	s6 := arr[1:4:5]
	fmt.Printf("len: %d, cap: %d \n", len(s6), cap(s6))

	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s7 := arr2[:5:6]
	fmt.Printf("len: %d, cap: %d \n", len(s7), cap(s7))

	// fmt.Println("slice ", appendSlice())
	fmt.Println("slice", copySliceV1())
	fmt.Println("slice", copySliceV2())
	fmt.Println("slice", copySliceV3())

	fmt.Println("slice", outSlice())

	// var arr [5]int = [5]int{1, 2, 3, 4, 5}
	// var s []int = arr[1:3]

	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	// fmt.Println(s[:cap(s)])

}

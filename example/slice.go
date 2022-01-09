package main

import "fmt"

func printSlice(arr []int) {
	fmt.Printf("arr len=%d cap=%d slice=%v\n", len(arr), cap(arr), arr)
}
func main() {
	fmt.Println("Hello World!")
	arr := []int{1, 2, 3}
	printSlice(arr)
	printSlice(arr[1:2])
	printSlice(arr)
	printSlice(append(arr, 1))

	var arr1 = make([]int, len(arr), cap(arr))
	copy(arr1, arr)
	printSlice(arr1)

}

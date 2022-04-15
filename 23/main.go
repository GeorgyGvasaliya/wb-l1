package main

import "fmt"

func DeleteByIndex(sl []int, i int) []int {
	return append(sl[:i], sl[i+1:]...)
}

func main() {
	nums := []int{2, 4, 6, 10, 12, 14, 16}
	fmt.Println(DeleteByIndex(nums, 1))
}

package main

import "fmt"

func BinSearch(sl []int, number int) int {
	l := 0
	r := len(sl) - 1

	for l <= r {
		mid := (l + r) / 2
		if sl[mid] < number {
			l = mid + 1
		} else if sl[mid] > number {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	//17. Бинарный поиск
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(BinSearch(nums, target))
}

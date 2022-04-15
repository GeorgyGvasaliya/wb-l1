package main

import "fmt"

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	ind := partition(nums, l, r)
	quickSort(nums, ind+1, r)
	quickSort(nums, l, ind-1)

}

func partition(nums []int, l, r int) int {
	p := l

	// nums[r] - pivot
	for i := l; i < r; i++ {
		if nums[i] <= nums[r] {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	nums[p], nums[r] = nums[r], nums[p]

	return p
}

func main() {
	//16. Быстрая сортировка
	nums := []int{1, 0, -4, 19, 22, 34, -22, 22, 8, 10, 1}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

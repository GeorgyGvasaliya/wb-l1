package main

import (
	"fmt"
)

func main() {
	var num int64 = 5
	fmt.Println("Before", num)
	fmt.Println("After", Set(num, 2, 0))
}

func Set(num, pos, val int64) int64 {
	if val == 1 {
		mask := int64(1) << pos
		return num | mask
	}
	return num &^ (int64(1) << pos)
}

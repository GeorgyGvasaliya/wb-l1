package main

import (
	"fmt"
)

func Groups(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)
	for _, v := range temps {
		groups[int(v)/10*10] = append(groups[int(v)/10*10], v)
	}
	return groups
}

func main() {
	nums := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := Groups(nums)
	fmt.Println(groups)
}

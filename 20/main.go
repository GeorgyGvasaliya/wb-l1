package main

import (
	"fmt"
	"strings"
)

func main() {
	//20. Разработать программу, которая переворачивает слова в строке.
	//Пример: «snow dog sun — sun dog snow».
	s := "snow dog sun"
	sSlice := strings.Fields(s)
	for i, j := 0, len(sSlice)-1; i < j; i, j = i+1, j-1 {
		sSlice[i], sSlice[j] = sSlice[j], sSlice[i]
	}
	fmt.Println(strings.Join(sSlice, " "))
}

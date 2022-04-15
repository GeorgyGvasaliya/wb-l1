package main

import "fmt"

func IsUnique(s string) bool {
	counter := make(map[rune]struct{})
	for _, ch := range s {
		if _, ok := counter[ch]; !ok {
			counter[ch] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

func main() {
	//26. Разработать программу, которая проверяет, что все символы в строке уникальные.
	//Функция проверки должна быть регистронезависимой.
	fmt.Println(IsUnique("deded"))
}

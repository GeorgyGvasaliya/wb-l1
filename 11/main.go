package main

import "fmt"

func Intersection(a, b []int) []int {
	m := make(map[int]int)
	c := make([]int, 0)

	for _, v := range a {
		m[v]++
	}

	for _, v := range b {
		if _, ok := m[v]; ok && m[v] > 0 {
			m[v]--
			c = append(c, v)
		}
	}
	return c
}

func main() {
	//Реализовать пересечение двух неупорядоченных множеств.
	num1 := []int{1, 1, 2, 2, 3, 4, 0, 1, 1, 1}
	num2 := []int{1, 1, 2, 3, 0, -4}

	fmt.Println(Intersection(num1, num2))

}

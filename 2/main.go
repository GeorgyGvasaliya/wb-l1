package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Pow(i int) {
	defer wg.Done()
	fmt.Println(i * i)
}

func main() {
	// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
	// и выведет их квадраты в stdout.
	sl := []int{2, 4, 6, 8, 10}

	for _, i := range sl {
		wg.Add(1)
		go Pow(i)
	}

	wg.Wait()
}

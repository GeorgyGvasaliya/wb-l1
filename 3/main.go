package main

import "fmt"

func main() {
	sl := []int{2, 4, 6, 8, 10}

	ch := make(chan int, len(sl))

	for i := 0; i < len(sl); i++ {
		go Sender(ch, sl[i])
	}

	sum := 0
	for i := 0; i < len(sl); i++ {
		sum += <-ch
	}
	close(ch)

	fmt.Println(sum)
}

func Sender(ch chan int, val int) {
	ch <- val * val
}

package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4}
	for v := range senderX2(sender(sl)) {
		fmt.Println(v)
	}
}

func sender(numbers []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range numbers {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func senderX2(ch <-chan int) <-chan int {
	doubleCh := make(chan int)
	go func() {
		for v := range ch {
			doubleCh <- v * 2
		}
		close(doubleCh)
	}()
	return doubleCh
}

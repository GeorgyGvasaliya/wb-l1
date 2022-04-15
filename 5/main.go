package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Разработать программу, которая будет последовательно отправлять значения в канал,
	//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
	ch := make(chan int)
	n := 4

	go receiver(ch)

	timeout := time.After(time.Duration(n) * time.Second)
	for {
		select {
		case <-timeout:
			fmt.Println("Time's up!")
			close(ch)
			goto next
		default:
			ch <- rand.Intn(10)
		}
	}

next:
	fmt.Println("end")
}

func receiver(ch chan int) {
	for v := range ch {
		fmt.Println("Got:", v)
	}
}

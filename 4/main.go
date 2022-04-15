package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(ch <-chan interface{}, id int) {
	for d := range ch {
		fmt.Println(d)
	}
}

func main() {
	var workersNum int

	fmt.Print("Number of workers: ")
	if _, err := fmt.Scan(&workersNum); err != nil {
		log.Fatalln("Bad number of workers value, program is shutting down")
	}

	done := make(chan os.Signal, 1)
	data := make(chan interface{})

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	for i := 0; i < workersNum; i++ {
		go worker(data, i)
	}
	for {
		select {
		case <-done:
			close(data)
			close(done)
			return
		default:
			data <- rand.Intn(1000)
			time.Sleep(time.Millisecond * 400)
		}
	}
}

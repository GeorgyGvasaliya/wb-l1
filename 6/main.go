package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func routineWithCancel(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}

func routineWithTimeOut(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}
func routineWithDeadline(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; i >= 0; i++ {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}

func routineWithCraftedCancelChan(cancel chan bool, wg *sync.WaitGroup) {
	for i := 0; i >= 0; i++ {
		select {
		case <-cancel:
			wg.Done()
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	cancelChan := make(chan bool)
	wg.Add(4)
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, _ := context.WithTimeout(context.Background(), 2*time.Second)
	ctx3, _ := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	go routineWithCancel(ctx1, &wg)
	go routineWithDeadline(ctx3, &wg)
	go routineWithTimeOut(ctx2, &wg)
	go routineWithCraftedCancelChan(cancelChan, &wg)
	go func() {
		time.Sleep(6 * time.Second)
		cancel1()
		time.Sleep(1 * time.Second)
		cancelChan <- true
	}()
	wg.Wait()
}

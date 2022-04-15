package main

import (
	"fmt"
	"sync"
)

type Incrementor struct {
	i int
	sync.Mutex
}

func (inc *Incrementor) DoIncrement() {
	inc.Lock()
	inc.i++
	inc.Unlock()
}

func main() {
	inc := Incrementor{}
	n := 1000
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			inc.DoIncrement()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(inc.i)
}

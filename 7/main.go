package main

import (
	"fmt"
	"sync"
)

type MyMap struct {
	table map[string]int
	sync.Mutex
}

func (m *MyMap) Read(key string) int {
	m.Lock()
	value := m.table[key]
	m.Unlock()
	return value
}

func (m *MyMap) Write(key string, value int) {
	m.Lock()
	m.table[key] = value
	m.Unlock()
}

func main() {
	//Реализовать конкурентную запись данных в map.
	mp := MyMap{table: make(map[string]int)}

	mp.Write("key1", 10)
	val := mp.Read("key1")
	fmt.Println(val)
}

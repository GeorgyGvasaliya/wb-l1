package main

import "fmt"

type Human struct {
	a, b int
	s    string
}

type Action struct {
	isEnd bool
	a     int
	Human
}

func main() {
	//Дана структура Human (с произвольным набором полей и методов).
	//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
	h := Human{a: 10, b: 11, s: "string"}
	a := Action{Human: h}
	fmt.Println(a.b)

}

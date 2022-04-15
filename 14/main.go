package main

import "fmt"

func CheckType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("it's an int")
	case string:
		fmt.Println("it's a string")
	case bool:
		fmt.Println("it's a bool")
	case chan int, chan string, chan bool:
		fmt.Println("it's a chan")
	}
}

func main() {
	ch := make(chan string)
	CheckType(3)
	CheckType(false)
	CheckType("string")
	CheckType(ch)
}

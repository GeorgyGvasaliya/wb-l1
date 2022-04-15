package main

import (
	"fmt"
	"math/big"
)

//func DoOperation(a, b string, foo func(x *big.Int, y *big.Int) *big.Int) *big.Int {
//	num1 := new(big.Int)
//	num2 := new(big.Int)
//
//	num1.SetString(a, 10)
//	num1.SetString(b, 10)
//
//	return foo(num1, num2)
//}

func main() {
	//22. Разработать программу, которая перемножает, делит, складывает,
	//вычитает две числовых переменных a,b, значение которых > 2^20.

	//a, b := "240000000000000000000000000000000000000000000000000000000000000", "2"
	////c, d, e, f, g, h := a, b, a, b, a, b
	//result1 := new(big.Int)
	//result2 := new(big.Int)
	//result3 := new(big.Int)
	////result4 := new(big.Int)
	//
	//s := DoOperation(a, b, result1.Add)
	//fmt.Printf("%T", s)
	//fmt.Println(DoOperation(a, b, result2.Sub))
	//fmt.Println(DoOperation(a, b, result3.Mul))
	////fmt.Println(DoOperation(a, b, result4.Div))

	x := big.NewInt(2931021)
	y := big.NewInt(2332121)

	fmt.Println(new(big.Int).Add(x, y))
	fmt.Println(new(big.Int).Sub(x, y))
	fmt.Println(new(big.Int).Mul(x, y))
	fmt.Println(new(big.Int).Div(x, y))
}

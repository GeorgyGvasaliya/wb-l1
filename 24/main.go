package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func CreatePoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func DistanceBetweenPoints(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func main() {
	p1 := CreatePoint(1, 2)
	p2 := CreatePoint(3, -5)
	fmt.Println(DistanceBetweenPoints(p1, p2))
}

package main

import (
	"fmt"
)

type coordinates interface {
	xaxis() int
	yaxis() int
}

type point2D struct {
	X int
	Y int
}

func (p point2D) xaxis() int {
	return p.X
}

func (p point2D) yaxis() int {
	return p.Y
}

func findCoordinates(c coordinates) {
	fmt.Println("X :", c.xaxis(), "Y :", c.yaxis())
}

type xpoint int

func (s xpoint) xaxis() int {
	return int(s)
}

func (s xpoint) yaxis() int {
	return 0
}

func main() {
	p := point2D{-3, 9}
	fmt.Println(p)
	findCoordinates(p)

	y := xpoint(10)
	findCoordinates(y)
}
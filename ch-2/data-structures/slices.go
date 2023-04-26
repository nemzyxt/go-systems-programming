// slices demo
package main

import (
	"fmt"
)

func change(x []int) {
	x[3] = -2
}

func printSlice(x []int) {
	for _, num := range x {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func main() {
	mySlice := []int{-1, 3, 5, 0, 7, 4}
	fmt.Println("Before change :")
	printSlice(mySlice)
	change(mySlice)
	fmt.Println("After change :")
	printSlice(mySlice)

	fmt.Printf("Before : cap:%d len:%d\n", cap(mySlice), len(mySlice))
	mySlice = append(mySlice, 100)
	fmt.Printf("After : cap:%d len:%d\n", cap(mySlice), len(mySlice))

	another := make([]int, 4)
	fmt.Println("New slice with 4 elements")
	printSlice(another)
}
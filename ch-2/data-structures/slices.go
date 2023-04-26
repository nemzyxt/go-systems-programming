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
}
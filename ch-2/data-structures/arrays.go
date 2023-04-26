// arrays demo
package main

import (
	"fmt"
)

func main() {
	arr := [4]int{1, 2, 5, -4}
	arr2D := [2][2]int{{1, 2}, {3, 4}}
	arr3D := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	fmt.Println(arr)
	fmt.Println(arr2D)
	fmt.Println(arr3D)
}
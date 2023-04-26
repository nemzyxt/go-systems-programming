package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{1, 2, 3, 4}
	amap := make(map[string]int)

	for i:=0; i<len(arr); i++ {
		fmt.Printf("%s ", strconv.Itoa(arr[i]))
		amap[strconv.Itoa(arr[i])] = arr[i]
	}
	fmt.Println()

	for k, v := range amap {
		fmt.Printf("%s : %d\n", k, v)
	}
}
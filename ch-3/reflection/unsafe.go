// running unsafe code in Go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value int64 = 5

	var p1 = &value
	var p2 = (*int32) (unsafe.Pointer(&value))

	fmt.Println("*p1 :", *p1)
	fmt.Println("*p2 :", *p2)
	*p1 = 312121321321213212
	fmt.Println("Value :", value)
	fmt.Println("*p2 :", *p2)
	*p1 = 31212132
	fmt.Println("Value :", value)
	fmt.Println("*p2 :", *p2)
}
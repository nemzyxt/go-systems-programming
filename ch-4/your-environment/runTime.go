// get runtime info using the runtime package
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Compiler :", runtime.Compiler)
	fmt.Println("Machine :", runtime.GOARCH)
	fmt.Println("Go Version :", runtime.Version())
	fmt.Println("Num. of Goroutines :", runtime.NumGoroutine())
}
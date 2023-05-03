// error accessing private items in package
package main

import (
	"fmt"
	"anotherPackage"
)

func main() {
	anotherPackage.Version()
	// fmt.Println(anotherPackage.version)
	fmt.Println(anotherPackage.Pi)
}
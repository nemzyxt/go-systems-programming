package anotherPackage

import "fmt"

const version = "1.1"
const Pi = "3.14159"

func init() {
	fmt.Println("The init function of anotherPackage")
}

func Version() {
	fmt.Println("The version of the package is :", version)
}
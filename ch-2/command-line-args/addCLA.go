// Sum all command line arguments
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	sum := 0
	for i := 0; i < len(arguments); i++ {
		n, _ := strconv.Atoi(arguments[i])
		sum += n
	}
	fmt.Println("Sum :", sum)
}
// Sum all command line arguments (with error handling)
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
		n, err := strconv.Atoi(arguments[i])
		if err == nil {
			sum += n
		} else {
			fmt.Println("Ignoring :", arguments[i])
		}		
	}
	fmt.Println("Sum :", sum)
}
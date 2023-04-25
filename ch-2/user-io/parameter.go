package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	minusI := false

	for i := 0; i < len(arguments); i++ {
		if strings.Compare(arguments[i], "-i") == 0 {
			minusI = true
			break
		}
	}

	if minusI {
		fmt.Println("Got the -i parameter")
		fmt.Print("y/n ? ")
		var ans string
		fmt.Scanln(&ans)
		fmt.Println("Ans :", ans)
	} else {
		fmt.Println("Didn't get the -i parameter")
	}
}
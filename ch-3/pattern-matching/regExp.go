// pattern-matching in Go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("Nemuel", "Nemuel Wainaina")
	fmt.Println(match)

	match, _ = regexp.MatchString("Wainaina", "Nemuel wainaina")
	fmt.Println(match)

	parse, err := regexp.Compile("[Nn]emuel")
	if err != nil {
		fmt.Println("Error compiling RE :", err)
	} else {
		fmt.Println(parse.MatchString("Nemuel Wainaina"))
		fmt.Println(parse.MatchString("nemuel Wainaina"))
		fmt.Println(parse.MatchString("N emuel Wainaina"))
		fmt.Println(parse.ReplaceAllString("Nemuel nemuel", "NEMUEL"))
	}
}
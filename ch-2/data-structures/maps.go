package main

import (
	"fmt"
)

func main() {
	amap := make(map[string]int)
	amap["mon"] = 0
	amap["tue"] = 1
	amap["wed"] = 2
	amap["thur"] = 3
	amap["fri"] = 4
	amap["sat"] = 5
	amap["sun"] = 6
	fmt.Printf("Sunday is the %dth day of the week\n", amap["sun"])

	_, ok := amap["Tuesday"]
	if ok {
		fmt.Println("The Tuesday key exists")
	} else {
		fmt.Println("The Tuesday key does not exist")
	}

	count := 0
	for key := range amap {
		count++
		fmt.Printf("%s ", key)
	}
	fmt.Println("\namap has", count, "elements")

	count = 0
	delete(amap, "thur")
	for range amap {
		count++
	}
	fmt.Println("amap now has", count, "elements")
}
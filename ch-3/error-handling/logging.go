// error logging in Go
package main

import (
	"log"
)

func main() {
	x := 1
	log.Println("log.Print() function :", x)
	x += 1
	log.Println("log.Print() function :", x)
	x += 1
	log.Panicln("log.Panic() function :", x)
	x += 1
	log.Println("log.Print() function :", x)
}
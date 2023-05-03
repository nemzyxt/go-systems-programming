// demo on reflection in Go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	type t1 int
	type t2 int

	x1 := t1(1)
	x2 := t2(1)
	x3 := 1

	st1 := reflect.ValueOf(&x1).Elem()
	st2 := reflect.ValueOf(&x2).Elem()
	st3 := reflect.ValueOf(&x3).Elem()

	type1 := st1.Type()
	type2 := st2.Type()
	type3 := st3.Type()

	fmt.Println("X1 Type :", type1)
	fmt.Println("X2 Type :", type2)
	fmt.Println("X3 Type :", type3)

	type aStructure struct {
		X uint
		Y float64
		Text string
	}
	x4 := aStructure{123, 3.14, "A Structure"}
	st4 := reflect.ValueOf(&x4).Elem()
	type4 := st4.Type()

	fmt.Println("The type of x4 is :", type4)
	fmt.Println("The fields of", type4, "are :")
	for i:=0; i<st4.NumField(); i++ {
		fmt.Printf("%d: Field name: %s ", i, type4.Field(i).Name)
		fmt.Printf("Type: %s ", st4.Field(i).Type())
		fmt.Printf("and Value: %v\n", st4.Field(i).Interface())
	}
}
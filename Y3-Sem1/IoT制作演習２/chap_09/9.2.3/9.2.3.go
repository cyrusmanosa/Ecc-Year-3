package main

import (
	"fmt"
	"reflect"
)

type point struct {
	 name string `color:"red" size:"12pt"`
	 x, y int64
}

func main() {
	var p1 point
	field := reflect.TypeOf(p1).Field(0)
	fmt.Println("color:", field.Tag.Get("color"))
	fmt.Println("size:", field.Tag.Get("size"))
}

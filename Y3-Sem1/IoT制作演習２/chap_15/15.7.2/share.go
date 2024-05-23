package main

import (
	"C"
	"fmt"
)

//export goprint
func goprint() {
	fmt.Println("Go Program")
}

func main() {}

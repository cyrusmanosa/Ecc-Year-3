//go:generate $GOPATH/bin/stringer -type DayOfWeek $GOFILE
package main

import "fmt"

type DayOfWeek int

const (
	Sunday DayOfWeek = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday, Saturday)
}

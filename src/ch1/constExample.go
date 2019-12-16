package main

import (
	"fmt"
	"unsafe"
)
const(
	Unknown = 0
    Female = 1
    Male = 2
)

type Weekday int
const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
func main() {
	const name = "hello woed"
	const (
		a = "123"
		b = 123
	)
	fmt.Println(name)
	fmt.Println(a)
	fmt.Println(b)	
	fmt.Println(Sunday)
	fmt.Println(unsafe.Sizeof(Male))
}

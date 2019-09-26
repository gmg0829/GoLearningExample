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
func main() {
	const name = "hello woed"
	const (
		a = "123"
		b = 123
	)
	fmt.Println(name)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(unsafe.Sizeof(Male))
}

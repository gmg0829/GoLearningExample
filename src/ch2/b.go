package main

import "fmt"

func main() {
	const name = "hello woed"
	const (
		a = "123"
		b = 123
	)
	fmt.Println(name)
	fmt.Println(a)
	fmt.Println(b)
}

package main

import "fmt"

/*
 引用类型 slice map channel
  使用make函数创建
*/

func makemap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}
func main() {
	m := makemap()
	fmt.Println(m["a"])
}

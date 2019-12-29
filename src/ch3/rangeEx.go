package main

import "fmt"

func  main(){
	hh := []int{1, 2, 3, 4, 5}
	for key, value := range hh {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}

	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// 通道
	c := make(chan int)

	go func() {

		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

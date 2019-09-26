package main

import "fmt"

func main()  {
	a:=[...]int{1,2}
	b:=[2]int{1,2}
	fmt.Println(a==b)
	months := [...]string{1: "January", 2: "February",3: "March",4: "April",5: "May"}
	Q2 := months[1:3]
	Q1 := months[:4]
	fmt.Println(Q2)
	fmt.Println(Q1)
}


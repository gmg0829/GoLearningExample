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

	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"

	for k, v := range team {
		fmt.Println(k, v)
	}

	// var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	// array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}

	// array = [4][2]int{1: {20, 21}, 3: {40, 41}}

	// 声明两个二维整型数组
	// var array1 [2][2]int
	// var array2 [2][2]int
	// // 为array2的每个元素赋值
	// array2[0][0] = 10
	// array2[0][1] = 20
	// array2[1][0] = 30
	// array2[1][1] = 40
	// // 将 array2 的值复制给 array1
	// array1 = array2

	// var a []int
	// a = append(a, 1) // 追加1个元素
	// a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	// copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1,slice2)

	
}


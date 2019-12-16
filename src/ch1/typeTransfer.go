package main

import(
	"fmt"
	"strconv"
)

func main() {
	// 整型转字符串
    num := 100
    str := strconv.Itoa(num)
	fmt.Printf("type:%T value:%#v\n", str, str)

	// 字符串转整型
	str1 := "110" //s100
    num1, err := strconv.Atoi(str1)
    if err != nil {
        fmt.Printf("%v 转换失败！", str1)
    } else {
        fmt.Printf("type:%T value:%#v\n", num1, num1)
	}
	// Parse 系列函数用于将字符串转换为指定类型的值，其中包括 ParseBool()、ParseFloat()、
	// ParseInt()、ParseUint()
	// ParseBool()  1、0、t、f、T、F、true、false、True、False、TRUE、FALSE，其它的值均返回错误

	str2 := "t"
	boo2, err := strconv.ParseBool(str2)
	fmt.Println(boo2)
	
}
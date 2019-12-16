/*
指针（pointer）在Go语言中可以被拆分为两个核心概念：
类型指针，允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无须拷贝数据，类型指针不能进行偏移和运算。
切片，由指向起始元素的原始指针、元素数量和容量组成。
*/
// 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

package main

import (
	"fmt"
	"flag"
)


// 定义命令行参数
var mode = flag.String("mode", "", "process mode")

func main() {
    var cat int = 1
	var str string = "banana"
	// 指针的值是带有0x十六进制前缀的一组数据
	fmt.Printf("%p %p", &cat, &str)
	
	 // 准备一个字符串类型
	 var house = "Malibu Point 10880, 90265"
	 // 对字符串取地址, ptr类型为*string
	 ptr := &house
	 // 打印ptr的类型
	 fmt.Printf("ptr type: %T\n", ptr)
	 // 打印ptr的指针地址
	 fmt.Printf("address: %p\n", ptr)
	 // 对指针进行取值操作
	 value := *ptr
	 // 取值后的类型
	 fmt.Printf("value type: %T\n", value)
	 // 指针取值后就是指向变量的值
	 fmt.Printf("value: %s\n", value)

	 // 准备两个变量, 赋值1和2
	 x, y := 1, 2
	 // 交换变量值
	 swap(&x, &y)
	 // 输出变量值
	 fmt.Println(x, y)

	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)


	// 创建指针的另一种方法——new() 函数

	str5 := new(string)
	*str5 = "Go语言教程"
	fmt.Println(*str5)
}


// 交换函数 a,b 指针类型
func swap(a, b *int) {
    // 取a指针的值, 赋给临时变量t
    t := *a
    // 取b指针的值, 赋给a指针指向的变量
    *a = *b
    // 将a指针的值赋给b指针指向的变量
    *b = t
}
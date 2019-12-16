package main
// Go语言的基本类型有：
// bool
// string
// int、int8、int16、int32、int64
// uint、uint8、uint16、uint32、uint64、uintptr
// byte // uint8 的别名
// rune // int32 的别名 代表一个 Unicode 码
// float32、float64
// complex64、complex128
// 匿名变量 “_”本身就是一个特殊的标识符，被称为空白标识符 

import ("fmt"
)

var (
	h int
	k bool
)
var c, d int = 1, 2

func GetData() (int, int) {
    return 100, 200
}

func main() {
	var x int //0
	fmt.Println(x)

	// var k bool //flase
	// var d,f string ="aaa","bbb"
	b, c := 100, "abc"
	fmt.Println(b, c)
	fmt.Println(&b)

	// var names []string
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!

	// const y = .71828 // 0.71828
	// const z = 1.     // 1

	// 使用内置的 complex 函数构建复数，并使用 real 和 imag 函数返回复数的实部和虚部
	// var x complex128 = complex(1, 2) // 1+2i
	// var y complex128 = complex(3, 4) // 3+4i
	// fmt.Println(x*y)                 // "(-5+10i)"
	// fmt.Println(real(x*y))           // "-5"
	// fmt.Println(imag(x*y))           // "10"


	k, _ := GetData()
    _, j := GetData()
    fmt.Println(k,j)

}



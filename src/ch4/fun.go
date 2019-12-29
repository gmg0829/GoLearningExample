package main

import "fmt"

func main()  {
	// const b=12.0;
	// fmt.Println(compute(b))
	a, b := namedRetValues()
    fmt.Println(a, b)

}

func compute(a float64) float64{
  return a*20
}

func typedTwoValues() (int, int) {
    return 1, 2
}

func namedRetValues() (a, b int) {
    a = 1
    b = 2
    return
}

func myfunc(args ...int) {
    for _, arg := range args {
        fmt.Println(arg)
    }
}
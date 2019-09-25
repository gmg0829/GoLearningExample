package main

import "fmt"

func main()  {
	const b=12.0;
	fmt.Println(compute(b))
}

func compute(a float64) float64{
  return a*20
}


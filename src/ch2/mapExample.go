package main
import "fmt"

func main(){
	var mapA map[string] int
	mapA=map[string] int{"one":1,"two":2}
	// mapCreated := make(map[string]float32)
	fmt.Println(mapA["mapA"])
	// 遍历
	for k, v := range mapA {
		fmt.Println(k, v)
	}

	//  删除
	delete(mapA, "one")

	
}
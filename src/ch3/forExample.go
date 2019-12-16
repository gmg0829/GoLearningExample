package main

import "fmt"

func main()  {
   sum:=0
   for index := 0; index < 10; index++ {
	   sum+=index
   }
   fmt.Println(sum)

   for {
      sum++
      if sum > 100 {
         break
      }
   }
}
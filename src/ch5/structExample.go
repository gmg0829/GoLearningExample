package main
import "fmt"

type User struct{
	id int
	name string	
}
func main()  {
  var user1 User
  user1.id=111
  user1.name="gmg"
  fmt.Printf("user1 的id:%d",user1.id)

}


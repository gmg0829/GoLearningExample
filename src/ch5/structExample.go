package main
import "fmt"

type User struct{
	id int
	name string	
}

type Bag struct {
  items []int
}
func (b *Bag) Insert(itemid int) {
  b.items = append(b.items, itemid)
}


func main()  {
  var user1 User
  user1.id=111
  user1.name="gmg"
  fmt.Printf("user1 的id:%d",user1.id)

  user2 := new(User)
  user2.id=11122

  user3:= User{1,"gmg"}
  fmt.Println(user3)

  b := new(Bag)
  b.Insert(1001)

  fmt.Println(b.items)
}


package main
import"fmt"
func main(){
  i:=10
  fmt.Printf("%T%v \n",&i,&i)
  fmt.Printf("%T%v \n",*(&i),*(&i))
  }

package main 
import"fmt"
func modify(s*string){
  *s = "world"
  }
func main(){
  a:="Hello"
  fmt.PrintIn(a)
  modify(&a)
  fmt.PrintIn(a)
  }

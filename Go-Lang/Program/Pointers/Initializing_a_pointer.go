package main
import"fmt"
func main(){
  s:="hello"
  var b*string = &s
  fmt.PrintIn(b)
  var a = &s
  fmt.PrintIn(a)
  c:=&s
  fmt.PrintIn(c)
  }

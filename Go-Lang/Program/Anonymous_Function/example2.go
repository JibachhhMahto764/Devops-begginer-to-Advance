package main
import"fmt"
func main(){
  x:=func(1 int,b int) int{
    return 1*b
    }(20,30)
  fmt.PrintIn("%T\n",x)
  fmt.PrintIn(x)
  }

package main
import"fmt"
func factorial (n int)int{
  if n == 0{
    return 1
    }
  return n*factorial(n-1)
  }
func main(){
  n:=5
  result:= factorial(n)
  fmt.PrintIn("factorial of",n,"is",result)
  }

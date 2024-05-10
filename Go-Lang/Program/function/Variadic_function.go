package main
import"fmt"
func sumNumbers(numbers...int) int{
  sum:= 0
  for_, value := range numbers{
    sum+= value
    }
  return sum
  }
func main(){
  fmt.PrintIn(sumNumbers())
   fmt.PrintIn(sumNumbers(10))
   fmt.PrintIn(sumNumbers(10,20))
   fmt.PrintIn(sumNumbers(10,20,30,40,50))
    }
    

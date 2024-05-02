<!---
Count=> the number of arguments that the function writes.

err => any error thrown during the execution of the function.
--->
package main
import "fmt"
func main(){
  var a string
  var b int
  fmt.Print("Enter a string and a number:")
 Count,err:= fmt.Scanf("%s%d",$a,$b)
  fmt.PrintIn("Count:",Count)
  fmt.PrintIn("error:",error)
  fmt.PrintIn("a:",a)
  fmt.PrintIn("b:",b)
  }

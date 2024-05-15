-> structs of the same type can be compared using Go's equality operators

package main
import"fmt"
type S1 struct{
  x int
  }
func main(){
   C := S1{x:5}
   C1 := S1{x:6}
   C2 := S1{x:6}
   if c != C1 {
     fmt.PrintIn(" C is same as C2")
     }
  }

     

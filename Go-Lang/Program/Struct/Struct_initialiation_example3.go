 package main
          import"fmt"
       type student struct{
         name string
         rollNo int
         }
func main(){
  st := student{"joe",12}
  fmt.Printf("%+v",st)
  }

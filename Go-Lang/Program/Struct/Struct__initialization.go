synatx =>
              var<variable_name><struct_name>
  such as ->     var s student


example;- 
          package main
          import"fmt"
       type student struct{
         name string
         rollNo int
         marks[] int
         grades map[string]int
         }
      func main(){
        var s student
        fmt.Printf("%+v",s)
        }

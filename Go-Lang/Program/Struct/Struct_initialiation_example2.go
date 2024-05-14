syntax =>  <variable_name> := new(<struct_name>)
st := new(student)




   package main
          import"fmt"
       type student struct{
         name string
         rollNo int
         marks[] int
         grades map[string]int
         }
      func main(){
        st := new(student)
        var s student
        fmt.Printf("%+v",s)
        }

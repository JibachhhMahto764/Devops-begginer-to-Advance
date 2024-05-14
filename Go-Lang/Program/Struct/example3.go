syntax=>
      <variable_name> := <struct_name>{
        <field_name> : <value>,
        <field_name> : <>value,
        }

example=> 

        package main
          import"fmt"
   type student struct{
         name string
         rollNo int
     } 
func main(){
  st := student{
    name:"joe"
    rollNo: 12,
    }
  fmt.Printf("%+v",st)
  }

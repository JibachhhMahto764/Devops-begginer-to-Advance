%t format specifier

reflect.Typeof => function from the reflect package

%T using....



package main
import "fmt"
func main(){
  var grades int = 42
  var message string = "hello-world"
  var isCheck bool = true
  var amount float32 = 5466.464

  fmt.printf("variable grades = "%v is of type %T \n", grades,grades)
  fmt.printf("variable grades = "%v is of type %T \n", message,message)
  fmt.printf("variable ischeck = "%v" is of type %T \n", ischeck,ischeck)
  fmt.printf("variable amount = %v is of type %T \n", amount,amount)
             }
  

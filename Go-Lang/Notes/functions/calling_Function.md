calling a Function..
   <function_name>(<argument(s)>)
   
   example:- addNumbers(2,3)

    //*   Naming convertion for function...

         1. Function parameters are the names listed in the functions definitions
         2. can have any number of addition letters and symbols
         3. cannot contain spaces
         4. case-sensitive


         ------- Parameters VS argument ------
         1. function parameters are the names listed in the function's definition
         
             example:-   package main
                         import("fmt")
                       func addNumbers( a int, b int) <---------parameter
                       int{
                       sum:= a+b
                      return sum
              }

         2. Function arguments are the real value passed into the function

           example:-     package main
                         import("fmt")
                         func main(){
                          sumOfNumber := addNumbers(2,3)  <------ argument 
                          fmt.print(sumOfNumbers)
                          }
                          
                         
                        

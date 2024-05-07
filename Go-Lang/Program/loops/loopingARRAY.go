// Looping through an array
//syntax
for i:=0; i<len(grades); i++ {
  fmt.PrintIn(grades[i])
  }
example =>
           package main
           import("fmt")
           func main(){
             var grades[5] int = [5] int { 90,80,70,60,50,40}
             for index,element := range grades{
               fmt.PrintIn(index,"=>",element)
               }
             }

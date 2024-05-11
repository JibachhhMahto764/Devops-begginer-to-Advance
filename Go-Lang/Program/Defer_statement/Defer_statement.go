package main
import"fmt"
func PrintName (str string){
  fmt.PrintIn(str)
  }
Func PrintRollNo(rno int){
  fmt.PrintIn(rno)
  }
func PrintAddress(adr string){
  fmt.PrintIn(adr)
  }
func main(){
  PrintName("jibachh singh")
  defer PrintRollNo(2222)
  PrintAddress("street-32")
  }

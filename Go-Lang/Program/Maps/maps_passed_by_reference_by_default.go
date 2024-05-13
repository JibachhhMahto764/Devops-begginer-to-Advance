package main
import"fmt"
func modify( m map[string] int){
  m["k"] = 75
  }
func main (){
  ascii_codes := make(map[string] int)
 ascii_codes["A"] = 65
   ascii_codes["F"] = 70
  fmt.PrintIn(ascii_codes)
  modify(ascii_codes)
  fmt.PrintIn(ascii_codes)
  }

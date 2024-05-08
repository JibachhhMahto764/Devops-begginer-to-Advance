package main
import("fmt")
func main(){
  codes := map[string] int {"en" :1."fr":2,"hi":3}
  value,found := codes["en"]
  fmt.PrintIn(found,value)
  value,found = codes["hh"]
  fmt.PrintIn(found,value)
  }

package main 
import("fmt")
func main(){
  var i int = 1000
  switch i { 
    case 10: 
       fmt.PrintIn("i is is 10")
    case 100,200:
    fmt.PrintIn("i is either 100 or 200")
    default:
    fmt.PrintIn("i is neither 0,1000 or 200")
    }
  }

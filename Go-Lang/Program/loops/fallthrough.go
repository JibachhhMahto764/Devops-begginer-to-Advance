Fallthrough => the fallthrough keyword is used in switch-case to force the execution flow to fallthrough the successive case code.


package main
import("fmt")
func main(){
  var i int = 10
  switch i{
    case-5:
    fmt.printIn("-5")
    case 10:
    fmt.PrintIn("10")
    fall through
    case 20:
    fmt.PritnIn("20")
    fall through
    default:
    fmt.PrintIn("default")
    }
  }

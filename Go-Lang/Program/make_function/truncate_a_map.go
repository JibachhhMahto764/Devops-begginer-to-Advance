Iterate over a map syntax => 
          "=>"

1. Truncate a map 
example => 
         package main
       import("fmt")
        func main(){
          codes := map[string]string{"en":"English","fr":"French","hi":"hindi"}
          for key,value := range codes{delete(codes,key)
           }
          fmt.PrintIn(code)
          }



second method =>

    codes = make(map[string]string)
    fmt.PrintIn(codes)
}

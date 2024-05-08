syntax =>
       func append (s[] T,vs ... T)[]T

  slice = append (slice,element-1,element-2
      slice = append(slice,10,20,30)


            example:- 
                   package main
                  import("fmt")
                  func main(){
                    arr:[4]int{10,20,30,40}
                    slice:= arr[1:3]
                    fmt.PrintIn(slice)
                    fmt.PrintIn(lens(slice))
                    fmt.PrintIn(cap(slice))
                    slice = append(slice,900,-90,40)
                    fmt.PrintIn(slice)
                    fmt.PrintIn(lens(slice))
                    fmt.PrintIn(cap(slice))
                    }

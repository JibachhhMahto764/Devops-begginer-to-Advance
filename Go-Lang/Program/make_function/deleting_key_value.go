deleting key-value-pair 
-> delete(map,key_name)


example => 
        package main
        import("fmt")
      func main(){
        codes := map[string]string{"en" :"English","fr":"French"}
        fmt.PrintIn(codes)
        delete(codes,"en")
        fmt.PrintIn(codes)
        }

-> An interface specifies a method set and is a powerful way to introduce modularity in Go. 
-> Interface is like a blueprint for a method set.
-> They describe all the methods of a method set by providing the function signature for each method.
-> They specify a set of methods but do not implement them.

syntax =>  type<interface_name>interface{
// method
}

such as => 
           type fixedDEposit interface{
           getRateofInterest() float64 clacReturn() float64
           }

implementing an interface

1. A type implements an interface by implementing its methods
2. The Go language interface are implemented implicity
3. And it does not have any specific keyword to implement an interface

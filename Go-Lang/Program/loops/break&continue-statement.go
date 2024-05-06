package main

import "fmt"

func main() {
    // Using break statement
    fmt.Println("Using break statement:")
    for i := 0; i < 10; i++ {
        if i == 5 {
            break // Exit the loop when i equals 5
        }
        fmt.Println(i)
    }

    // Using continue statement
    fmt.Println("\nUsing continue statement:")
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue // Skip even numbers
        }
        fmt.Println(i)
    }
}

package main

import "fmt"

func bitwiseOperators(a, b int) {
    // Bitwise AND
    fmt.Println("Bitwise AND (a & b):", a & b)

    // Bitwise OR
    fmt.Println("Bitwise OR (a | b):", a | b)

    // Bitwise XOR
    fmt.Println("Bitwise XOR (a ^ b):", a ^ b)

    // Bitwise NOT (Complement)
    fmt.Println("Bitwise NOT (^a):", ^a)
    fmt.Println("Bitwise NOT (^b):", ^b)

    // Bitwise Left Shift
    shiftAmount := 2
    fmt.Println("Bitwise Left Shift (a <<", shiftAmount, "):", a << shiftAmount)
    fmt.Println("Bitwise Left Shift (b <<", shiftAmount, "):", b << shiftAmount)

    // Bitwise Right Shift
    fmt.Println("Bitwise Right Shift (a >>", shiftAmount, "):", a >> shiftAmount)
    fmt.Println("Bitwise Right Shift (b >>", shiftAmount, "):", b >> shiftAmount)
}

func main() {
    // Sample values
    num1 := 15 // Binary: 1111
    num2 := 7  // Binary: 0111

    fmt.Println("Using num1 =", num1, "and num2 =", num2)
    bitwiseOperators(num1, num2)
}


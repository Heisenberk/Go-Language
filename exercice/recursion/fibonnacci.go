package main

import "fmt"

func fibonacci(n int) int {
    if n <= 0 {
        return 1
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    fmt.Println(fibonacci(7))
}
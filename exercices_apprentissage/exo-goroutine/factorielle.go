package main
import (
    "fmt"
    "sync"
    "time"
)
var wg sync.WaitGroup


var sum_goroutine int

var sum int



func factorial_goroutine() {
    for i := 0; i < 1000000; i++ {
        sum_goroutine += i
    }
    wg.Done()
}


func factorial() {
    for i := 0; i < 1000000; i++ {
        sum += i
    }

}
func main() {
    fmt.Println("################################# Avec goroutine #############################")
    start := time.Now()
    wg.Add(1)
    go factorial_goroutine()
    wg.Wait()
    fmt.Println("temps d'exécution", time.Since(start))
    fmt.Println("résultat", sum_goroutine)
    
    fmt.Println("################################# Sans goroutine #############################")
    start2 := time.Now()
    factorial()
    fmt.Println("temps d'exécution", time.Since(start2))
    fmt.Println("résultat", sum)
}
package main

import "fmt"
import "point"

func main() {
    p := point.Point{X : 3, Y: 2}
    fmt.Println(point.PrintPoint(p))
}
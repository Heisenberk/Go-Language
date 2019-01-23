package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["h1"]=1
    m["h2"]=2
    fmt.Println("map :",m)
    delete(m,"h1");
    fmt.Println("delete map",m)
}
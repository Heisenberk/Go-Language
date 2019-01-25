package main

import "fmt"


func main() {
 secondes := 3954
var heure int = secondes/3600 
var minute int = (secondes % 3600)/60
var seconde int = secondes%60

fmt.Println("heure :",heure)
fmt.Println("minute :",minute)
fmt.Println("seconde :", seconde);

}
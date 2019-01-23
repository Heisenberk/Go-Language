package main

import "fmt"

type Equipement interface {
porter() int8
}

type Soldat struct {
speed,endurance int8
}

func (s Soldat) porter() int8 {
return s.speed - 2
}
func mesure(e Equipement){
fmt.Println(e)
fmt.Println("speed",e.porter())

}
func main() {
s := Soldat{speed : 5, endurance : 3}
mesure(s)

}
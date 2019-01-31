package main

import "fmt"
import "strings"

func main() {
	tab := []string{"clement", "mehdi", "TER", "master"}
	chaine := strings.Join(tab, ", ")
	fmt.Println(chaine)
}
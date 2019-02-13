package main

import "fmt"
import "os"
import "strconv"


func main() {
	if secondes, err := strconv.Atoi(os.Args[1]); err == nil {

		fmt.Println("Entrée",secondes,"secondes")
		
		var heure int = secondes/3600 
		var minute int = (secondes % 3600)/60
		var seconde int = secondes%60

		fmt.Println("heure :",heure)
		fmt.Println("minute :",minute)
		fmt.Println("seconde :", seconde);
	}

}
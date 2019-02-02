package main

import "fmt"

func  CalculDiv(u int) int {
	var somme int
	for i:=1 ; i<u ; i++ {
		if u%i==0{

			somme=somme+i
		}

	}

	somme=somme +u
	return somme; 

}

func main() {
	var n int
	var m int
	n = 220 
	m = 284
	if n<0 || m<0 {
		fmt.Println("n et m ne sont pas amis")
	}else{

		if CalculDiv(n)==CalculDiv(m) {
			fmt.Println("n et m  sont amis")
		}else{
			fmt.Println("n et m ne sont pas amis")
		}


	}

	
   
}
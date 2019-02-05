
package main

import (
		"fmt"

)

type Polynome struct {
    Tab [NMAX]int 
    Degre int
}

	const NMAX int=50

func Init_polynome() Polynome{

	var P Polynome

	for i:=0 ; i<NMAX ;i++ {

		P.Tab[i]=0
	}

	return P

}

func Inserer_polynome(coef int ,degre int,P Polynome) Polynome{

	P.Tab[degre]=coef

	return P
}

func Affiche_polynome(P Polynome){

	var u int 

	

}

func main(){

	var P Polynome

	P=Init_polynome()

	fmt.Println(P)

	fmt.Println("------------------------")

	P=Inserer_polynome(2,0,P)

	P=Inserer_polynome(3,1,P)

	P=Inserer_polynome(6,2,P)

	fmt.Println(P)
}

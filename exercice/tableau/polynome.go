
package main

import (
		"fmt"

)

type Polynome struct {
    Tab []int 
    Degre int
}

	

func Init_polynome() Polynome{

	var P Polynome

	P.Tab=make([]int,50)

	for i:=0 ; i<len(P.Tab) ;i++ {

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

	u=0

		for u<len(P.Tab){

			if P.Tab[u]!=0 && u==0{

				fmt.Print(P.Tab[0])

			}

			if P.Tab[u]!=0 {

				fmt.Print(" + ",P.Tab[u],"^",u)

			}

			u++
		}

	fmt.Println()
}

func Somme_polynome(P1 Polynome,P2 Polynome) Polynome{

	var P3 Polynome

	var etat1 int 

	var etat2 int 

	if len(P1.Tab)>len(P2.Tab) {

		P3.Tab=make([]int,len(P1.Tab))

	}else{

		P3.Tab=make([]int,len(P2.Tab))

	}


	etat1=0

	etat2=0

	for i:=0 ;i<len(P3.Tab);i++ {

		if i<len(P1.Tab){


			if P1.Tab[i]!=0 {

			etat1=1

			}

		}
		if i<len(P2.Tab) {

			if i<len(P2.Tab) {

				if P2.Tab[i]!=0{

					etat2=2

				}

			}

		}

	}

	if etat1 == 0 && etat2 == 0{return P3}
	if etat1 == 1 && etat2 == 0{return P1}
	if etat1 == 0 && etat2 == 1{return P2} 

	if len(P1.Tab) > len (P2.Tab){

		for i:=0; i<len(P1.Tab); i++{

			if i<len(P2.Tab){

				P3.Tab[i]=P1.Tab[i] + P2.Tab[i]

			}else{

				P3.Tab[i]=P1.Tab[i]
			}

		}

	}else{

		for i:=0; i<len(P2.Tab); i++{

			if i<len(P1.Tab){

				P3.Tab[i]=P1.Tab[i] + P2.Tab[i]

			}else{

				P3.Tab[i]=P2.Tab[i]
			}

		}
	}	

	return P3

}


func main(){

	var P Polynome

	var P3 Polynome

	P=Init_polynome()

	P3.Tab=make([]int,55)



	Affiche_polynome(P)

	fmt.Println("------------------------")

	P3=Inserer_polynome(9,0,P3)

	P=Inserer_polynome(2,5,P)

	P=Inserer_polynome(3,1,P)

	P=Inserer_polynome(6,2,P)

	P4:=Somme_polynome(P,P3)

	fmt.Println(P4)

	Affiche_polynome(P)




	//var n int  



	
}

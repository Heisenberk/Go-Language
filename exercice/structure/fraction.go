package main

import "fmt"

type Fraction struct {
	Denominateur int
	Numerateur int
}

func Pgcd(a int,b int) int{

	m:=b

	if a<b {

		m=a

	}

	for a % m !=0 || b % m !=0 {

		m--

	}
	
	return m
}

func Init_frac(num int,den int) Fraction{

	div :=Pgcd(num,den)

	var resultat Fraction

	resultat.Numerateur=num/div

	resultat.Denominateur=den/div

	return resultat
	
}

func Afficher_frac(f Fraction) {

	fmt.Println(f.Numerateur,"/",f.Denominateur)
	
}

func Add_frac(f1 Fraction,f2 Fraction) Fraction{

	return Init_frac(f1.Numerateur   * f2.Denominateur + f2.Numerateur * f1.Denominateur,f1.Denominateur * f2.Denominateur)

}

func Mult_frac(f1 Fraction,f2 Fraction) Fraction{

	return Init_frac(f1.Numerateur   * f2.Numerateur,f1.Denominateur * f2.Denominateur)

}

func Mult_scal_frac(f1 Fraction,scalaire int) Fraction{

	return Init_frac(f1.Numerateur * scalaire, f1.Denominateur)

}

func main() {

	f1 := Init_frac(10,18)

	f2 := Init_frac(20,18)

	Afficher_frac(f1)

	Afficher_frac(f2)

	Afficher_frac(Add_frac(f1,f2))

	
}

package main

import "fmt"
import "strconv"
import "os"

type Fraction struct {
	Numerateur int
	Denominateur int
} 

func PGCD(a int, b int) int{
	m:=b
	if a<b {
		m=a
	}
	for a % m !=0 || b % m !=0 {
		m--
	}
	return m
}

func InitFrac(num int, den int) Fraction{
	div :=PGCD(num,den)
	var resultat Fraction
	resultat.Numerateur=num/div
	resultat.Denominateur=den/div
	return resultat
}

func AfficherFrac(f Fraction) {
	fmt.Println(f.Numerateur,"/",f.Denominateur)
}

// A/B + C/D = (AD/BD)+(CD/BD) = (AD+CD)/BD
func AddFrac(f1 Fraction,f2 Fraction) Fraction{
	return InitFrac(f1.Numerateur * f2.Denominateur + f2.Numerateur * f1.Denominateur, f1.Denominateur * f2.Denominateur)
}

// A/B - C/D = (AD/BD)-(CD/BD) = (AD-CD)/BD
func SubFrac(f1 Fraction,f2 Fraction) Fraction{
	return InitFrac(f1.Numerateur * f2.Denominateur - f2.Numerateur * f1.Denominateur, f1.Denominateur * f2.Denominateur)
}

// A/B * C/D = AC/BD
func MultFrac(f1 Fraction,f2 Fraction) Fraction{
	return InitFrac(f1.Numerateur * f2.Numerateur, f1.Denominateur * f2.Denominateur)
}

// A/B / C/D = A/B * D/C = AD/BC
func DivFrac(f1 Fraction,f2 Fraction) Fraction{
	return InitFrac(f1.Numerateur * f2.Numerateur, f1.Denominateur * f2.Denominateur)
}


func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("Aucun argument. ")
	}else if len(arg) == 1 {
		fmt.Println("Manque un argument. ")
	}else {
		val1, err1 := strconv.Atoi(arg[0])
		val2, err2 := strconv.Atoi(arg[1])
		if err1 != nil {
			fmt.Println("Argument 1 invalide. ")
		}else if err2 != nil {
			fmt.Println("Argument 2 invalide. ")
		}else {
			var f Fraction
			f = InitFrac(val1, val2)
			AfficherFrac(f)
		}
	}
}
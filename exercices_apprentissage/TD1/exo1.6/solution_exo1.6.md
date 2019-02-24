# TD1 Exercice 6 : Solution 

1. La structure possède deux champs `Numerateur` et `Denominateur`. 
```go
type Fraction struct {
	Numerateur int
	Denominateur int
} 
```

2. Avant toute chose, il faut rendre les fractions irréductibles. Pour cela, créer une fonction supplémentaire qui calcule le PGCD du numérateur et du dénominateur : 
```go
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
```
Ensuite, créer la fonction `InitFraction` et toutes les autres fonctions de calcul qui appeleront `InitFraction`. Enfin, créer la fonction d'affichage. 

```go
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
```

Pour avoir le code complet : [exo1.6](exo1.6.zip)
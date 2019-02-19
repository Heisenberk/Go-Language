
# TD1 : Bases du langage Go

But : Apprendre l'environnement et manipuler les différentes bases du langage de programmation Go. 

Difficulté : Facile

Connaissances à avoir : Variables, Conditions, Fonctions,  Bibliothèques standards, Ligne de commande et Arguments, Structures.

Durée du TD : 1h15 - 1h30

## Exercice 1 : Installation

Avant de programmer en langage Go, il faut installer l'environnement. Pour cela, aller sur la documentation officielle de [Golang](https://golang.org/doc/install).

1. Installer l'environnement pour programmer en langage Go. 

2. Vérifier que la compilation et l'exécution d'un programme basique "Hello, world!" fonctionne correctement. 

Solution: 

1. Aller sur le site de Golang pour récupérer le fichier `.tar.gz` (pour Linux). Par exemple, la dernière version disponible est `go1.11.5.linux-amd64.tar.gz`. Puis, taper la commande suivante pour installer l'environnement de `Go`: 
```shell
tar -C /usr/local -xzf go1.11.5.linux-amd64.tar.gz
```
Ensuite, aller dans le fichier `/home/user/.profile` et ajouter à la fin du fichier la ligne `PATH="/usr/local/go/bin:$PATH"`. 
Puis, pour installer les commandes `go`, taper la commande `sudo apt-get install golang-go`.
Enfin, taper la commande `go env` et vérifier le chemin pour `$GOPATH`. Ce chemin va correspondre au répertoire où il faudra programmer en langage Go. 

2. Dans le dossier `$GOPATH/src/`, créer le dossier `hello` dans lequel il faudra créer le fichier `hello.go`. 

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world! \n")
}
```
Pour compiler, il suffit de taper `go build hello`, puis, pour exécuter, faire `./hello`. 

Pour avoir le code complet : [exo1.1](exo1.1.zip)

## Exercice 2 : Conversions

Ecrire un programme qui convertit un temps donné en secondes, minutes et secondes (avec l'accord des pluriels). Le calcul doit se faire dans une fonction et doit retourner, par exemple, `3620 correspond à 1 heure, 0 minute et 20 secondes`. 

Note : Utiliser la bibliothèque standard `strconv` pour convertir un entier en string. 

Solution : 

Créer une fonction qui prend un entier en paramètre et qui renvoie un string. Calculer dans cette fonction les différentes valeurs des secondes, des minutes et des heures. Utiliser les conditions pour savoir si il y a une ou plusieurs heures, minutes et secondes. Convertir les entiers en string avec `strconv.Itoa(...)`. Retourner la concaténation de tous les éléments et appeler cette fonction dans le main. 

```go
package main

import "fmt"
import "strconv"

func conversion (val int) string {
	var heure int = val / 3600 
	var minute int = (val % 3600) / 60
	var seconde int = val % 60

	var phrase string = " correspond à "

	var heure_str string;
	var minute_str string;
	var seconde_str string;

	if heure <= 1 {
		heure_str = " heure, "
	}else{
		heure_str = " heures, "
	} 

	if minute <= 1 {
		minute_str = " minute, "
	}else{
		minute_str = " minutes, "
	} 

	if seconde <= 1 {
		seconde_str = " seconde. "
	}else{
		seconde_str = " secondes. "
	} 
	
	return strconv.Itoa(val) + phrase + strconv.Itoa(heure) + heure_str + strconv.Itoa(minute) + minute_str + strconv.Itoa(seconde) + seconde_str;
}


func main() {
	fmt.Println(conversion(3620))
}
```

Pour avoir le code complet : [exo1.2](exo1.2.zip)


## Exercice 3 : Récursion

Il est possible de faire des appels récursifs en langage Go. Fibonacci et Factorielle sont deux fonctions illustrant la récursivité. 

1. Ecrire un programme récursif qui permet de calculer la `n` iéme valeur de la suite de Fibonacci. 

2. Ecrire un programme récursif qui permet de calculer la factorielle de `n`.

L'utilisateur du programme doit pouvoir entrer en paramètre d'arguments le nombre à calculer. 

Solution : 

Exercice de Fibonacci :

```go
package main

import "fmt"
import "strconv"
import "os"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("Aucun argument. ")
	}else {
		val, err := strconv.Atoi(arg[0])
		if err != nil {
			fmt.Println("Argument invalide. ")
		}else if val < 0 {
			fmt.Println("Impossible de calculer la factorielle d'un nombre négatif. ")
		}else {
			fmt.Println(fibonacci(val))
		}
	}
}
```

Exercice de Factorielle : 

```go
package main

import "fmt"
import "strconv"
import "os"

func factorielle (n int) int {
    if n == 0 {
        return 1
    }
    return n * factorielle(n-1)
}

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("Aucun argument. ")
	}else {
		val, err := strconv.Atoi(arg[0])
		if err != nil {
			fmt.Println("Argument invalide. ")
		}else if val < 0 {
			fmt.Println("Impossible de calculer la factorielle d'un nombre négatif. ")
		}else {
			fmt.Println(factorielle(val))
		}
	}
}
```

Pour avoir le code complet : [exo1.3](exo1.3.zip)


## Exercice 4 : Primalité

Ecrire un fonction `primalite` qui teste de la manière la plus naïve possible la primalité d'un nombre, et une fonction `interpreter` qui retourne la chaîne de caractères affichée sur le terminal. L'utilisateur du programme doit pouvoir entrer en paramètre le nombre à tester. 

Solution : 

Tester de manière naïve la primalité d'un nombre consiste à tester s'il est divisible par un nombre plus petit que lui et différent de 2. De cette façon, créer la fonction `primalite`. De plus, il faut interpréter le résultat de cette dernière par le biais de `interpreter`. 

```go
package main

import "fmt"
import "os"
import "strconv"

func primalite (val int) int {
	if val <= 1 {
		return 0
	}else if val == 2 {
		return 1
	}else {
		for i := 2 ; i < val ; i++ {
			if val % i == 0 {
				return 0
			}
		}
	}
	return 1
}

func interpreter (val int) string {
	var test int = primalite(val)

	if test == 0 {
		return "n n'est pas premier. "
	}else{
		return "n est premier. "
	}
}

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("Aucun argument. ")
	}else {
		val, err := strconv.Atoi(arg[0])
		if err != nil {
			fmt.Println("Argument invalide");
		}else {
			fmt.Println(interpreter(val))
		}
	}
}
```

Pour avoir le code complet : [exo1.4](exo1.4.zip)

## Exercice 5 : Nombres Amicaux

Soit `n` et `m`, deux entiers positifs. `n` et `m` sont dit amis si la somme de tous les diviseurs de `n` (sauf `n` lui-même) est égale à `m` et si la somme de tous les diviseur de `m` (sauf `m` lui-même) est égale à `n`. 
Ecrire un programme qui teste si deux entiers sont des nombres amis ou non. L'utilisateur du programme doit pouvoir entrer en paramètre le nombre à tester. 

Solution : 

Créer une fonction `sommeDiviseurs` qui calcule la somme des diviseurs de l'entier en entrée. Créer une fonction `amis` qui teste si les sommes des diviseurs des deux entiers en entrée sont égales : si c'est le cas alors ils sont amis. La fonction `interpreter` permet de renvoyer une chaîne de caractères pour y être affichée dans le terminal. 

```go
package main

import "fmt"
import "strconv"
import "os"

func sommeDiviseurs (val int) int {
	var somme int
	for i:=1 ; i<val ; i++ {
		if val%i==0{
			somme=somme+i
		}
	}

	somme=somme +val
	return somme; 
}

func amis (val1 int, val2 int) int {
	if val1<0 || val2<0 {
		return 0
	}else{
		if sommeDiviseurs(val1)== sommeDiviseurs(val2) {
			return 1
		}else{
			return 0
		}
	}
}

func interpreter (val1 int, val2 int) string{
	if amis (val1, val2) == 0 {
		return "n et m ne sont pas amis. "
	}else {
		return "n et m sont amis. "
	}
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
			fmt.Println(interpreter(val1, val2))
		}
	}
}
```

Pour avoir le code complet : [exo1.5](exo1.5.zip)

## Exercice 6 : Fraction

1. Définir une structure de données `Fraction` permettant de représenter en forme fractionnaire un nombre flottant. De plus, la fraction doit toujours être irréductible. 
2. Ecrire les fonctions qui permettent d'initialiser, d'afficher une fraction, d'additionner, de soustraire, de multiplier et de diviser 2 fractions. 

Solution : 

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

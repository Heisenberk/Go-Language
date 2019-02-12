
# TD1 : Bases du langage Go

But : Apprendre l'environnement et manipuler les différentes bases du langage de programmation Go. 
Difficulté : Facile
Connaissances à avoir : Variables, Conditions, Fonctions,  Bibliothèques standards, Ligne de commande et Arguments, 

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

## Exercice 3 : Primalité

Ecrire un fonction `primalite` qui teste de la manière la plus naïve possible la primalité d'un nombre, et une fonction `interpreter` qui retourne la chaîne de caractères affichée sur le terminal. L'utilisateur du programme doit pouvoir entrer en paramètre le nombre à tester. 

Solution : 

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


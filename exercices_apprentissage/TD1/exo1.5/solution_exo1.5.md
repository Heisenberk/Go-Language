# TD1 Exercice 5 : Solution 

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